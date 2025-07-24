package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log/slog"
)

func Load() (*Config, *SQLConfig, error) {
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using environment variables")
	}

	config := new(Config)
	if err := env.Parse(config); err != nil {
		return nil, nil, err
	}

	sqlConfig := new(SQLConfig)
	if err := env.Parse(sqlConfig); err != nil {
		return nil, nil, err
	}

	return config, sqlConfig, nil
}
