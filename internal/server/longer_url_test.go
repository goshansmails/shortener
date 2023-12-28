package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goshansmails/shortener/internal/store/mockstore"
	"github.com/stretchr/testify/require"
)

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

	s := New(Settings{Store: store})

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, test.path, nil)
			w := httptest.NewRecorder()

			s.LongerURL(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if test.ok {
				require.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
				require.Equal(t, test.desiredLink, resp.Header.Get("Location"))
			} else {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			}
		})
	}
}
