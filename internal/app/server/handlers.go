package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/goshansmails/shortener/internal/app/store"
)

const linkFormat = "http://localhost:8080/%d"

func wrappedGetIDHandler(s store.Store) func(resp http.ResponseWriter, req *http.Request) {

	return func(resp http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		urlToSave := string(body)
		id, err := s.GetID(urlToSave)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.Header().Set("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusCreated)
		resp.Write([]byte(fmt.Sprintf(linkFormat, id)))
	}
}

func wrappedGetURLHandler(s store.Store) func(resp http.ResponseWriter, req *http.Request) {

	return func(resp http.ResponseWriter, req *http.Request) {

		path := strings.Trim(req.URL.Path, "/")
		id, err := strconv.Atoi(path)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		url, err := s.GetURL(id)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		resp.Header().Set("Location", url)
		resp.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func mainHandlerWrapped(s *Server) func(resp http.ResponseWriter, req *http.Request) {

	getIDHandler := wrappedGetIDHandler(s.store)
	getURLHandler := wrappedGetURLHandler(s.store)

	return func(resp http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodPost:
			getIDHandler(resp, req)
		case http.MethodGet:
			getURLHandler(resp, req)
		default:
			resp.WriteHeader(http.StatusBadRequest)
		}
	}
}
