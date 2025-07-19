package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"minemetrics_golang/internal/config"
	"minemetrics_golang/internal/handler"
)

func Run(cfg *config.Config, logger *slog.Logger) error {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Handlers
	metricsHandler := handler.NewMetricHandler(logger)

	router.Post("/metrics", metricsHandler.HandlePost)

	return http.ListenAndServe(":"+cfg.Port, router)
}
