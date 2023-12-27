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
	r := chi.NewRouter()

	r.Post("/", s.ShortenURL)
	r.Get("/*", s.LongerURL)

	return http.ListenAndServe(s.addr, r)
}
