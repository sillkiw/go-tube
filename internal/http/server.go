package httpserver

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	videosapi "github.com/sillkiw/gotube/internal/http/api/videos"
)

type Server struct {
	l      *slog.Logger
	router *chi.Mux
	vh     *videosapi.VideosHandler
}

func New(l *slog.Logger, vh *videosapi.VideosHandler) *Server {
	s := &Server{
		l:      l,
		router: chi.NewRouter(),
		vh:     vh,
	}
	s.routes()
	return s
}
