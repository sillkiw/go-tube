package http

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	videosapi "github.com/sillkiw/gotube/internal/http/api/videos"
	mvLogger "github.com/sillkiw/gotube/internal/http/middleware"
)

func NewRouter(logger *slog.Logger, vh *videosapi.VideosHandler) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(mvLogger.New(logger))
	router.Use(middleware.URLFormat)

	router.Mount("/api/videos", vh.NewRouter())

	return router
}
