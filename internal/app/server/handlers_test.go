package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/goshansmails/shortener/internal/app/store/mockstore"
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

	getIDHandler := wrappedGetIDHandler(mockstore.New())

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, test.path, strings.NewReader("https://ya.ru"))
			w := httptest.NewRecorder()
			getIDHandler(w, req)

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

func TestGetURLHandler(t *testing.T) {

	store := mockstore.New()

	store.AddPair("https://ya.ru", 1)
	store.AddPair("https://www.iana.org", 2)

	type TestCase struct {
		name        string
		path        string
		ok          bool
		desiredLink string
	}

	tests := []TestCase{
		{
			name:        "success #1",
			path:        "/1/",
			ok:          true,
			desiredLink: "https://ya.ru",
		},
		{
			name:        "success #2",
			path:        "/2",
			ok:          true,
			desiredLink: "https://www.iana.org",
		},
		{
			name: "unknown id",
			path: "/3",
			ok:   false,
		},
		{
			name: "unparsable path",
			path: "/a/1/",
			ok:   false,
		},
	}

	getURLHandler := wrappedGetURLHandler(store)

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, test.path, nil)
			w := httptest.NewRecorder()
			getURLHandler(w, req)

			resp := w.Result()
			if test.ok {
				require.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
				require.Equal(t, test.desiredLink, resp.Header.Get("Location"))
			} else {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			}
		})
	}
}
