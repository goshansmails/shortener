package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ShortenURL(longURL string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, c.hostPort, strings.NewReader(longURL))
	if err != nil {
		return "", fmt.Errorf("can't create request: %w", err)
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("status code = %d; can't read body: %w", resp.StatusCode, err)
		}
		return "", fmt.Errorf("status code: %d; body: %s", resp.StatusCode, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("can't read body: %w", err)
	}

	return string(body), nil
}
