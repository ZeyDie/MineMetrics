package responses

import (
	"encoding/json"
	"minemetrics_golang/internal/models"
	"net/http"
)

func JSON(responseWriter http.ResponseWriter, status int, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)

	json.NewEncoder(responseWriter).Encode(data)
}

func Success(responseWriter http.ResponseWriter, message string) {
	Response(
		responseWriter,
		http.StatusOK,
		models.Response{
			Status:  "success",
			Message: message,
		},
	)
}

func Error(responseWriter http.ResponseWriter, statusCode int, message string) {
	Response(
		responseWriter,
		statusCode,
		models.Response{
			Status:  "error",
			Message: message,
		},
	)
}

func Response(responseWriter http.ResponseWriter, statusCode int, reponse models.Response) {
	JSON(responseWriter, statusCode, reponse)
}

func DecodeJSON(request *http.Request, data interface{}) error {
	return json.NewDecoder(request.Body).Decode(data)
}
