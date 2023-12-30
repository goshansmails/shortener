package client

import (
	"net/http"
)

type Client struct {
	hostPort string
	client   *http.Client
}

func New(hostPort string) *Client {
	return &Client{
		hostPort: hostPort,
		client: &http.Client{
			// no redirect
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}
