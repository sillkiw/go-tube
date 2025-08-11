package main

import (
	"flag"
	"gotube/internal/app"
	"gotube/internal/config"
	"gotube/internal/cookie"
	"gotube/internal/templates"
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
	configPath = "./config.yaml"
	// sessionIDLength = 32
)

func main() {
	configPath := flag.String("config", configPath, "Path to configuration file")
	flag.Parse()

	// Load Config
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Define log type
	var handler slog.Handler
	if cfg.Mode == "production" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}
	logger := slog.New(handler)                                                   // Application logger
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // Server error logger

	// Parse YAML user file
	usersPath := cfg.Auth.UsersFilePath
	users, err := user.LoadUsersFromFile(usersPath)
	if err != nil {
		logger.Error("failed to load users",
			slog.String("path", usersPath),
			slog.Any("err", err),
		)
		os.Exit(1)
	}
	logger.Info("users loaded",
		slog.String("path", usersPath),
		slog.Int("count", len(users)),
	)

	// Generates 3 secret keys for key rotation
	cookie.InitializeKeys(3)

	// Initialize a new template cache
	templatePath := cfg.UI.TemplateDir
	templateCache, err := templates.CreateTemplateCache(templatePath)
	if err != nil {
		logger.Error("failed to load templates",
			slog.String("path", templatePath),
			slog.Any("err", err),
		)
		os.Exit(1)
	}
	logger.Info("templates loaded",
		slog.String("path", templatePath),
		slog.Int("count", len(templateCache)),
	)

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
	app := app.NewApplication(logger, cfg, templateCache, users)

	// go resetVideoUploadedCounter()
	if cfg.Server.TLS.Enabled {
		go func() {
			tlsSrv := &http.Server{
				Addr:     cfg.Server.TLS.BindAddress + ":" + cfg.Server.TLS.Port,
				ErrorLog: errorLog,
				Handler:  app.Routes(),
			}
			if err := tlsSrv.ListenAndServeTLS(cfg.Server.TLS.Cert, cfg.Server.TLS.Key); err != nil {
				logger.Error("TLS server failed", "err", err)
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
			errorLog.Fatal(err)
		}
	}

}
