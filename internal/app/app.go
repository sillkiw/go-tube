package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"time"

	"github.com/sillkiw/gotube/internal/config"
	"github.com/sillkiw/gotube/internal/filestore/disk"
	transport "github.com/sillkiw/gotube/internal/http"
	videosapi "github.com/sillkiw/gotube/internal/http/api/videos"
	videosvalidation "github.com/sillkiw/gotube/internal/http/api/videos/validation"
	"github.com/sillkiw/gotube/internal/httpserver"
	"github.com/sillkiw/gotube/internal/storage/postgres"
	"github.com/sillkiw/gotube/internal/videos"
)

type App struct {
	l      *slog.Logger
	cfg    config.Config
	db     io.Closer
	server *httpserver.Server
}

func New(logger *slog.Logger, errorLog *log.Logger, cfg config.Config) (*App, error) {
	const op = "app.New"

	a := &App{
		l:   logger,
		cfg: cfg,
	}

	storage, err := postgres.New(a.cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	a.db = storage

	fileStore := disk.New(cfg.Video.DashPath, cfg.Video.RawPath)

	videosServ := videos.New(storage, fileStore)
	validator := videosvalidation.New(cfg.Validation)
	videosHandler := videosapi.New(logger, videosServ, validator)

	mainHandler := transport.NewRouter(logger, videosHandler)

	server := httpserver.New(errorLog, mainHandler, cfg.Server)
	a.server = server

	return a, nil
}

func (a *App) Run() error {
	return a.server.Start()
}

func (a *App) Close() {
	if a.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := a.server.Shutdown(ctx); err != nil {
			a.l.Error("failed to shutdown server", slog.Any("err", err))
		}
	}

	if a.db != nil {
		if err := a.db.Close(); err != nil {
			a.l.Error("failed to close db", slog.Any("err", err))
		}
	}
}
