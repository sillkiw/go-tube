package videosapi

import (
	"log/slog"

	valid "github.com/sillkiw/gotube/internal/validation"
	"github.com/sillkiw/gotube/internal/videos"
)

type VideosHandler struct {
	logger    *slog.Logger
	svc       *videos.Service
	validator valid.Validator
}

func New(logger *slog.Logger, svc *videos.Service, v valid.Validator) *VideosHandler {
	return &VideosHandler{logger: logger, svc: svc, validator: v}
}
