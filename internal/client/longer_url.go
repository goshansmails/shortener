package client

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) LongerURL(shortURL string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, shortURL, nil)
	if err != nil {
		return "", fmt.Errorf("can't create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return "", fmt.Errorf("short url not found or request is invalid")
	}

	if resp.StatusCode != http.StatusTemporaryRedirect {
		return "", errors.New("no redirect status code found")
	}

	longURL := resp.Header.Get("Location")
	if longURL == "" {
		return "", errors.New("no location found")
	}

	return longURL, nil
}
