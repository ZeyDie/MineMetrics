package main

import (
	"log/slog"
	"minemetrics_golang/internal/config"
	"minemetrics_golang/internal/database"
	"minemetrics_golang/internal/server"
	"minemetrics_golang/internal/systemlog"
	"os"
)

var Version = "dev"

func main() {
	systemlog.Init(Version)

	slog.Info("Starting application", "version", Version)

	config, sqlConfig, err := config.Load()

	if err != nil {
		slog.Error(
			"Failed to load config",
			"error",
			err,
		)
		os.Exit(1)
	}

	slog.Info(
		"Loaded configuration",
		"port", config.Port,
	)

	database.NewDB(sqlConfig)

	if err := server.Run(config); err != nil {
		slog.Error(
			"Server error",
			"error",
			err,
		)
		os.Exit(1)
	}
}
