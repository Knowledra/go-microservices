package callAPI

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func CallAPI(method, url string, payload interface{}) ([]byte, error) {
	client := &http.Client{}

	var body io.Reader
	if payload != nil {
		jsonData, _ := json.Marshal(payload)
		body = bytes.NewBuffer(jsonData)
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
