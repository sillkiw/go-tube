package app

import (
	"gotube/internal/config"
	"gotube/internal/video"
	"html/template"
	"log/slog"
)

type Application struct {
	logger        *slog.Logger
	config        *config.Config
	templateCache map[string]*template.Template
	videoSrv      *video.Service
}

func NewApplication(logger *slog.Logger, config *config.Config, templateCache map[string]*template.Template) *Application {
	return &Application{
		logger:        logger,
		config:        config,
		templateCache: templateCache,
		videoSrv:      video.NewService(logger, config),
	}
}
