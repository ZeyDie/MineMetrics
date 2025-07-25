package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"minemetrics_golang/internal/config"
	"minemetrics_golang/internal/handler"
)

func Run(config *config.Config) error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json"))

	clientHandler := handler.NewClientHandler()
	serverHandler := handler.NewServerHandler()

	router.Post("/client", clientHandler.HandlePost)
	router.Post("/server", serverHandler.HandlePost)

	return http.ListenAndServe(":"+config.Port, router)
}
