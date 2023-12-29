package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goshansmails/shortener/internal/store/mockstore"
	"github.com/stretchr/testify/require"
)

func TestLongerURL(t *testing.T) {
	t.Skip()

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

	s := New(Settings{Store: store})

	server := httptest.NewServer(s.router)
	defer server.Close()

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodGet,
				server.URL+test.path,
				nil,
			)
			require.NoError(t, err)

			t.Log("aaaaa", server.URL+test.path)

			client := &http.Client{}

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
