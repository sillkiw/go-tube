package app

import (
	"gotube/internal/config"
	"gotube/internal/user"
	"gotube/internal/video"
	"html/template"
	"log/slog"
)

type Application struct {
	logger        *slog.Logger
	config        *config.Config
	templateCache map[string]*template.Template
	videoSrv      *video.Service
	users         []user.User
}

func NewApplication(logger *slog.Logger, config *config.Config, templateCache map[string]*template.Template, users []user.User) *Application {
	return &Application{
		logger:        logger,
		config:        config,
		templateCache: templateCache,
		videoSrv:      video.NewService(logger, config),
		users:         users,
	}
}
