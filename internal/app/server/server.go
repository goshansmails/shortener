package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/goshansmails/shortener/internal/app/store"
)

type Server struct {
	addr    string
	baseURL string
	store   store.Store
}

type Settings struct {
	Addr    string
	BaseURL string
	Store   store.Store
}

func New(settings Settings) *Server {
	return &Server{
		addr:    settings.Addr,
		baseURL: settings.BaseURL,
		store:   settings.Store,
	}
}

func (s *Server) Run() error {

	getIDHandler := wrappedGetIDHandler(s.store, s.baseURL)
	getURLHandler := wrappedGetURLHandler(s.store)

	r := chi.NewRouter()
	r.Post("/", getIDHandler)
	r.Get("/*", getURLHandler)

	return http.ListenAndServe(s.addr, r)
}
