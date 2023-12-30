package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goshansmails/shortener/internal/store/mockstore"
	"github.com/stretchr/testify/require"
)

func TestLongerURL(t *testing.T) {
	type TestCase struct {
		name        string
		path        string
		ok          bool
		desiredLink string
	}

	tests := []TestCase{
		{
			name:        "success #1",
			path:        "/1",
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
	}

	store := mockstore.New()
	store.AddPair("https://ya.ru", 1)
	store.AddPair("https://www.iana.org", 2)

	router := newRouter(newServer(Settings{
		Store: store,
	}))

	server := httptest.NewServer(router)
	defer server.Close()

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodGet,
				server.URL+test.path,
				nil,
			)
			require.NoError(t, err)

			// no redirect!
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			if test.ok {
				require.Equal(t, resp.StatusCode, http.StatusTemporaryRedirect)
				require.Equal(t, resp.Header.Get("Location"), test.desiredLink)
			} else {
				require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			}
		})
	}
}
