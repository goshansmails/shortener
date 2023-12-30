package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/goshansmails/shortener/internal/store/mockstore"
	"github.com/stretchr/testify/require"
)

var linkRegExp = regexp.MustCompile("http://localhost:8080/[0-9A-Za-z]+")

func TestShortenURL(t *testing.T) {
	type TestCase struct {
		name         string
		urlToShorten string
	}

	tests := []TestCase{
		{
			name:         "sample 1",
			urlToShorten: "https://ya.ru",
		},
		{
			name:         "sample 2",
			urlToShorten: "https://www.iana.org",
		},
		{
			name:         "sample 3",
			urlToShorten: "https://google.com",
		},
	}

	router := newRouter(newServer(Settings{
		Store:   mockstore.New(),
		BaseURL: "http://localhost:8080",
	}))

	server := httptest.NewServer(router)
	defer server.Close()

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPost,
				server.URL,
				strings.NewReader("https://ya.ru"),
			)
			require.NoError(t, err)

			client := &http.Client{}

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			require.Equal(t, http.StatusCreated, resp.StatusCode)
			require.Equal(t, "text/plain", resp.Header.Get("Content-Type"))

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.True(t, linkRegExp.MatchString(string(body)))
		})
	}
}
