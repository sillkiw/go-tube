package http

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	l      *slog.Logger
	router *chi.Mux
}

func New(l *slog.Logger) *Server {
	s := &Server{
		l:      l,
		router: chi.NewRouter(),
	}
	s.routes()
	return s
}
