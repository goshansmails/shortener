package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/goshansmails/shortener/internal/store/storeutils"
)

func (s *Server) LongerURL(resp http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := s.store.GetURL(id)
	if err != nil {
		fmt.Println("can't get url:", err)
		resp.WriteHeader(http.StatusBadRequest)
		if errors.Is(err, storeutils.ErrNotFound) {
			_, _ = resp.Write([]byte("short URL not found"))
		}
		return
	}

	resp.Header().Set("Location", url)
	resp.WriteHeader(http.StatusTemporaryRedirect)
}
