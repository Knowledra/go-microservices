package callAPI

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Singleton HTTP client with connection pooling and timeout.
var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

func CallAPI(c context.Context, method, url string, payload any) ([]byte, error) {
	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(c, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	var reqID string
	if c, ok := c.(*gin.Context); ok {
		reqID = c.GetString("X-Request-Id")
	} else if val, ok := c.Value("X-Request-Id").(string); ok {
		reqID = val
	}
	if reqID != "" {
		req.Header.Set("X-Request-Id", reqID)
	}

	if ginCtx, ok := c.(*gin.Context); ok {
		if authHeader := ginCtx.GetHeader("Authorization"); authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// APIRequest describes a single outbound request for parallel execution.
type APIRequest struct {
	Method  string
	URL     string
	Payload any
}

// APIResult holds the response (or error) for one parallel request.
type APIResult struct {
	Body []byte
	Err  error
}

// CallAPIsParallel fires multiple API requests concurrently and returns
// results in the same order as the input slice.
func CallAPIsParallel(c context.Context, requests []APIRequest) []APIResult {
	results := make([]APIResult, len(requests))
	var wg sync.WaitGroup

	var derivedCtx context.Context = c
	if ginCtx, ok := c.(*gin.Context); ok {
		derivedCtx = ginCtx.Copy()
	}

	for i, r := range requests {
		wg.Add(1)
		go func(idx int, req APIRequest) {
			defer wg.Done()
			body, err := CallAPI(derivedCtx, req.Method, req.URL, req.Payload)
			results[idx] = APIResult{Body: body, Err: err}
		}(i, r)
	}

	wg.Wait()
	return results
}
