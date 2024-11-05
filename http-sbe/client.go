package httpsbe

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient}
}

func (c *Client) DoRequest(ctx context.Context, method, url string, jsonData interface{}) ([]byte, error) {
	req, err := createRequest(ctx, method, url, jsonData)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bb, nil
}

func createRequest(ctx context.Context, method, url string, jsonData interface{}) (*http.Request, error) {
	var buf io.Reader
	if jsonData != nil {
		body, err := json.Marshal(jsonData)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/sbe")
	req.Header.Set("X-MBX-SBE", "2:0")

	return req, nil
}
