package app

import (
	"log/slog"

	"github.com/sillkiw/gotube/internal/config"
)

type App struct {
	log *slog.Logger
	cfg config.Config
}

func New(logger *slog.Logger, config config.Config) (*App, error) {
	a := &App{
		log: logger,
		cfg: config,
	}

	if err := a.initDB(); err != nil {

	}
}

func (a *App) initDB() {

}
