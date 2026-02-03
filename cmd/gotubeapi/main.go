package main

import (
	"gotube/internal/app"
	"gotube/internal/config"
	"gotube/internal/user"
	"log"
	"log/slog"
	"net/http"
	"os"
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

	// Parse YAML user file
	users := user.MustLoad(cfg.Auth.UsersFilePath)

	// d, err := time.ParseDuration(cfg.Video.DeleteOld.CheckInterval)
	// if err != nil {
	// 	fmt.Println("Error parsing CheckOldEvery from config.yaml. Using default value (1h)", err)
	// 	d = time.Hour
	// }
	// checkOldEvery = d
	// if cfg.Video.DeleteOld.Enabled {
	// 	go deleteOLD()
	// }

	// Initialize a new instance of application containing the dependencies
	app := app.NewApplication(logger, cfg, users)

	// go resetVideoUploadedCounter()
	if cfg.Server.TLS.Enabled {
		go func() {
			tlsSrv := &http.Server{
				Addr:     cfg.Server.TLS.BindAddress + ":" + cfg.Server.TLS.Port,
				ErrorLog: errorLog,
				Handler:  app.Routes(),
			}
			if err := tlsSrv.ListenAndServeTLS(cfg.Server.TLS.Cert, cfg.Server.TLS.Key); err != nil {
				logger.Error("failed to start TLS server",
					slog.Any("err", err),
				)
			}
		}()
	}
	if cfg.Server.HTTP.Enabled {
		httpSrv := &http.Server{
			Addr:     cfg.Server.HTTP.BindAddress + ":" + cfg.Server.HTTP.Port,
			ErrorLog: errorLog,
			Handler:  app.Routes(),
		}
		if err := httpSrv.ListenAndServe(); err != nil {
			logger.Error("failed to start server",
				slog.Any("err", err),
			)
		}
	}
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
