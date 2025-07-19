package responses

import (
	"encoding/json"
	"minemetrics_golang/internal/model"
	"net/http"
)

func JSON(responseWriter http.ResponseWriter, status int, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)

	json.NewEncoder(responseWriter).Encode(data)
}

func Success(responseWriter http.ResponseWriter, message string) {
	JSON(responseWriter, http.StatusOK, model.Response{
		Status:  "success",
		Message: message,
	})
}

func Error(responseWriter http.ResponseWriter, statusCode int, message string) {
	JSON(responseWriter, statusCode, model.Response{
		Status:  "error",
		Message: message,
	})
}

func DecodeJSON(request *http.Request, data interface{}) error {
	return json.NewDecoder(request.Body).Decode(data)
}
