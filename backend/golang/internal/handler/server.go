package handler

import (
	"minemetrics_golang/internal/database/metrics"
	"minemetrics_golang/internal/request"
	"minemetrics_golang/pkg/responses"
	"net/http"
)

type ServerHandler struct{}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{}
}

func (serverHandler *ServerHandler) HandlePost(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	request, err := DecodeMetric[request.ServerRequest](responseWriter, httpRequest)

	if err != nil {
		return
	}

	err = metrics.InsertServerData(request)

	if err != nil {
		responses.Error(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	responses.Success(responseWriter, "")
}
