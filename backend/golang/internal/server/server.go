package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"minemetrics_golang/internal/config"
	"minemetrics_golang/internal/handler"
)

func Run(cfg *config.Config) error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json"))

	metricsHandler := handler.NewMetricHandler()

	router.Post("/metrics", metricsHandler.HandlePost)

	return http.ListenAndServe(":"+cfg.Port, router)
}
