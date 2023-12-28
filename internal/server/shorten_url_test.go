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

func TestGetIDHandler(t *testing.T) {
	type TestCase struct {
		name string
		path string
		ok   bool
	}

	tests := []TestCase{
		{
			name: "success",
			path: "/",
			ok:   true,
		},
		{
			name: "bad path",
			path: "/a/b/c",
			ok:   false,
		},
	}

	s := New(Settings{
		Store:   mockstore.New(),
		BaseURL: "http://localhost:8080",
	})

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, test.path, strings.NewReader("https://ya.ru"))
			w := httptest.NewRecorder()

			s.ShortenURL(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if test.ok {
				require.Equal(t, http.StatusCreated, resp.StatusCode)
				require.Equal(t, "text/plain", resp.Header.Get("Content-Type"))

				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				require.True(t, linkRegExp.MatchString(string(body)))

			} else {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			}
		})
	}
}
