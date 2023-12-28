package server

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) LongerURL(resp http.ResponseWriter, req *http.Request) {
	path := strings.Trim(req.URL.Path, "/")
	id, err := strconv.Atoi(path)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := s.store.GetURL(id)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.Header().Set("Location", url)
	resp.WriteHeader(http.StatusTemporaryRedirect)
}
