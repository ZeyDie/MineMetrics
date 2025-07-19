package handler

import (
	"log/slog"
	"net/http"

	"minemetrics_golang/internal/model"
	"minemetrics_golang/pkg/responses"
)

type MetricHandler struct {
	logger *slog.Logger
}

func NewMetricHandler(logger *slog.Logger) *MetricHandler {
	return &MetricHandler{logger: logger}
}

func (metricHandler *MetricHandler) HandlePost(responseWriter http.ResponseWriter, request *http.Request) {
	var req model.ClientRequest

	if err := responses.DecodeJSON(request, &req); err != nil {
		responses.Error(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}

	metricHandler.logger.Info("POST", "data", req)

	responses.Success(responseWriter, "")
}
