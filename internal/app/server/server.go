package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/goshansmails/shortener/internal/app/store"
)

type Server struct {
	store store.Store
}

func New(store store.Store) *Server {
	return &Server{
		store: store,
	}
}

const linkFormat = "http://localhost:8080/%d"

func wrappedGetIDHandler(s *Server) func(resp http.ResponseWriter, req *http.Request) {

	return func(resp http.ResponseWriter, req *http.Request) {

		body, err := io.ReadAll(req.Body)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		urlToSave := string(body)
		id, err := s.store.GetID(urlToSave)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.Header().Set("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusCreated)
		resp.Write([]byte(fmt.Sprintf(linkFormat, id)))
	}
}

func wrappedGetURLHandler(s *Server) func(resp http.ResponseWriter, req *http.Request) {

	return func(resp http.ResponseWriter, req *http.Request) {

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
}

func mainHandlerWrapped(s *Server) func(resp http.ResponseWriter, req *http.Request) {

	getIDHandler := wrappedGetIDHandler(s)
	getURLHandler := wrappedGetURLHandler(s)

	return func(resp http.ResponseWriter, req *http.Request) {

		if req.URL.Path == "/" {
			if req.Method == http.MethodPost {
				getIDHandler(resp, req)
			} else {
				http.Error(resp, "", http.StatusBadRequest)
			}
			return
		}

		getURLHandler(resp, req)
	}

}

func (s *Server) Run() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandlerWrapped(s))

	return http.ListenAndServe(`:8080`, mux)
}
