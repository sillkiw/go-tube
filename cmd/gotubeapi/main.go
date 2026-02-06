package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/sillkiw/gotube/internal/app"
	"github.com/sillkiw/gotube/internal/config"
	"github.com/sillkiw/gotube/internal/user"
)

// var (
// 	checkOldEvery = time.Hour //wait time before recheck  file deletion policies
// )

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	// sessionIDLength = 32
)

func main() {
	cfg := config.MustLoad()

	logger, errorLog := setupLogger(cfg.Env)

	users := user.MustLoad(cfg.Auth.UsersFilePath)

	app := app.New(logger, cfg, users)

	logger.Error("server stopped")
}

func setupLogger(env string) (*slog.Logger, *log.Logger) {
	var logger *slog.Logger
	var errorLog *log.Logger

	var handler slog.Handler
	switch env {
	case envLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		handler = slog.NewTextHandler(os.Stdout, nil)
	}

	logger = slog.New(handler)
	errorLog = slog.NewLogLogger(handler, slog.LevelError)

	return logger, errorLog
}
