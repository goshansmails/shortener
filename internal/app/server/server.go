package server

import (
	"net/http"

	"github.com/go-chi/chi"
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

	getIDHandler := wrappedGetIDHandler(s.store)
	getURLHandler := wrappedGetURLHandler(s.store)

	r := chi.NewRouter()
	r.Post("/", getIDHandler)
	r.Get("/*", getURLHandler)

	return http.ListenAndServe(`:8080`, r)
}
