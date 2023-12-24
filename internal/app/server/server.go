package server

import (
	"net/http"

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

func (s *Server) Run() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandlerWrapped(s))

	return http.ListenAndServe(`:8080`, mux)
}
