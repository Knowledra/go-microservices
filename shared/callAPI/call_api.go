package callAPI

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallAPI(ctx context.Context, method, url string, payload interface{}) ([]byte, error) {
	client := &http.Client{}

	var body io.Reader
	if payload != nil {
		jsonData, _ := json.Marshal(payload)
		body = bytes.NewBuffer(jsonData)
	}

	req, _ := http.NewRequestWithContext(ctx, method, url, body)
	req.Header.Set("Content-Type", "application/json")

	var reqID string
	if c, ok := ctx.(*gin.Context); ok {
		reqID = c.GetString("X-Request-Id")
	} else if val, ok := ctx.Value("X-Request-Id").(string); ok {
		reqID = val
	}
	if reqID != "" {
		req.Header.Set("X-Request-Id", reqID)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
