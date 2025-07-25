package handler

import (
	"log/slog"
	"minemetrics_golang/pkg/responses"
	"net/http"
)

func DecodeMetric[T any](responseWriter http.ResponseWriter, httpRequest *http.Request) (T, error) {
	var request T

	if err := responses.DecodeJSON(httpRequest, &request); err != nil {
		slog.Error("POST", "err", err)
		responses.Error(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return request, err
	}

	slog.Info("POST", "request", request)

	return request, nil
}
