package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log/slog"
)

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using environment variables")
	}

	cfg := new(Config)

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
