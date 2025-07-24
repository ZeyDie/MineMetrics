package handler

import (
	"log/slog"
	"net/http"

	"minemetrics_golang/internal/model"
	"minemetrics_golang/pkg/responses"
)

type MetricHandler struct {
}

func NewMetricHandler() *MetricHandler {
	return &MetricHandler{}
}

func (metricHandler *MetricHandler) HandlePost(responseWriter http.ResponseWriter, request *http.Request) {
	var clientRequest model.ClientRequest

	if err := responses.DecodeJSON(request, &clientRequest); err != nil {
		responses.Error(responseWriter, http.StatusBadRequest, "Invalid request payload")
		slog.Error("POST", "err", err)
		return
	}

	slog.Info("POST", "clientRequest", clientRequest)

	responses.Success(responseWriter, "")
}
