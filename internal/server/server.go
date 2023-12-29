package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	addr    string
	baseURL string
	store   Store
}

type Settings struct {
	Addr    string
	BaseURL string
	Store   Store
}

type Store interface {
	GetID(url string) (int, error)
	GetURL(id int) (string, error)
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
	r.Get("/{id}", s.LongerURL)

	return http.ListenAndServe(s.addr, r)
}
