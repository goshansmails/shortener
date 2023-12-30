package server

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) ShortenURL(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	urlToSave := string(body)
	id, err := s.store.GetID(urlToSave)
	if err != nil {
		fmt.Println("can't get id:", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "text/plain")
	resp.WriteHeader(http.StatusCreated)
	_, _ = resp.Write([]byte(fmt.Sprintf("%s/%d", s.baseURL, id)))
}
