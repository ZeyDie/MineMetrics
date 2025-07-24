package handler

import (
	"log/slog"
	"minemetrics_golang/internal/database"
	"minemetrics_golang/internal/database/entity"
	"minemetrics_golang/internal/models/dto"
	"net/http"

	"minemetrics_golang/internal/models"
	"minemetrics_golang/pkg/responses"
)

type MetricHandler struct {
}

func NewMetricHandler() *MetricHandler {
	return &MetricHandler{}
}

func (metricHandler *MetricHandler) HandlePost(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	var request models.Request

	if err := responses.DecodeJSON(httpRequest, &request); err != nil {
		responses.Error(responseWriter, http.StatusBadRequest, "Invalid request payload")
		slog.Error("POST", "err", err)
		return
	}

	slog.Info("POST", "request", request)

	clientDTO, err := dto.ClientToDTO(&request)
	if err != nil {
		slog.Error("Can't converting to client DTO", "err", err)
		return
	}

	clientEntity, err := entity.ClientDTOToEntity(clientDTO)
	if err != nil {
		slog.Error("Can't converting to client entity", "err", err)
		return
	}

	database.GetConnection().Create(&clientEntity)

	responses.Success(responseWriter, "")
}
