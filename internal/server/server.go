package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	addr    string
	baseURL string
	store   Store
	router  *chi.Mux
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
	s := &Server{
		addr:    settings.Addr,
		baseURL: settings.BaseURL,
		store:   settings.Store,
	}

	r := chi.NewRouter()
	r.Post("/", s.ShortenURL)
	r.Get("/{id}", s.LongerURL)

	s.router = r

	return s
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.addr, s.router)
}
