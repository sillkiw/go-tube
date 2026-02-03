package app

import (
	"gotube/internal/config"
	"gotube/internal/user"
	"gotube/internal/video"
	"log/slog"
)

type Application struct {
	log      *slog.Logger
	cfg      *config.Config
	videoSrv *video.Service
	users    []user.User
}

func NewApplication(logger *slog.Logger, config *config.Config, users []user.User) *Application {
	return &Application{
		log:      logger,
		cfg:      config,
		videoSrv: video.NewService(logger, config),
		users:    users,
	}
}
