package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	baseURL string
	store   Store
}

type Settings struct {
	BaseURL string
	Store   Store
}

type Store interface {
	GetID(url string) (int, error)
	GetURL(id int) (string, error)
}

func newServer(settings Settings) *Server {
	return &Server{
		baseURL: settings.BaseURL,
		store:   settings.Store,
	}
}

func newRouter(s *Server) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.ShortenURL)
	r.Get("/{id}", s.LongerURL)

	return r
}

func Run(addr string, settings Settings) error {
	s := newServer(settings)
	r := newRouter(s)

	return http.ListenAndServe(addr, r)
}
