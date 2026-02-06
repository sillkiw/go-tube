package videosapi

import "github.com/go-chi/chi/v5"

func (vh *VideosHandler) NewRouter() *chi.Mux {
	videosRouter := chi.NewRouter()
	videosRouter.Post("/", vh.create)

	return videosRouter
}
