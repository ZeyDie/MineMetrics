package main

import (
	"log/slog"
	"minemetrics_golang/internal/config"
	"minemetrics_golang/internal/server"
	"os"
)

var Version = "dev"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	logger.Info("Starting application", "version", Version)

	cfg, err := config.Load()

	if err != nil {
		logger.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Info("Loaded configuration",
		"port", cfg.Port,
		"env", cfg.Env,
	)

	if err := server.Run(cfg, logger); err != nil {
		logger.Error("Server error", "error", err)
		os.Exit(1)
	}
}
