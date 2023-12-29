package server

import (
	"fmt"
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
		fmt.Println("can't get url:", err)
		resp.WriteHeader(http.StatusBadRequest)
		_, _ = resp.Write([]byte("short URL not found"))
		return
	}

	resp.Header().Set("Location", url)
	resp.WriteHeader(http.StatusTemporaryRedirect)
}
