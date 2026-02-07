package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/sillkiw/gotube/internal/app"
	"github.com/sillkiw/gotube/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	configPath := flag.String("config", "./config/config.yaml", "Path to configuration file")
	flag.Parse()
	cfg := config.MustLoad(*configPath)

	logger, errorLog := setupLogger(cfg.Env)

	// users := user.MustLoad(cfg.Auth.UsersFilePath)

	app, err := app.New(logger, errorLog, cfg)
	if err != nil {
		logger.Error("failed to init app",
			slog.Any("err", err),
		)
		return
	}
	defer app.Close()

	go func() {
		if err := app.Run(); err != nil {
			logger.Error("server error", slog.Any("err", err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
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
