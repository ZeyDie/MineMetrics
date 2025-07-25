package handler

import (
	"minemetrics_golang/internal/database/metrics"
	"minemetrics_golang/internal/request"
	"minemetrics_golang/pkg/responses"
	"net/http"
)

type ClientHandler struct {
}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{}
}

func (clientHandler *ClientHandler) HandlePost(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	request, err := DecodeMetric[request.ClientRequest](responseWriter, httpRequest)

	if err != nil {
		return
	}

	err = metrics.InsertClientData(request)

	if err != nil {
		responses.Error(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	responses.Success(responseWriter, "")
}
