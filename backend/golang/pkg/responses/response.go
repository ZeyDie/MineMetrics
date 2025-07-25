package responses

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func JSON(responseWriter http.ResponseWriter, status int, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)

	json.NewEncoder(responseWriter).Encode(data)
}

func Success(responseWriter http.ResponseWriter, message string) {
	Response(
		responseWriter,
		http.StatusOK,
		StatusResponse{
			Status:  "success",
			Message: message,
		},
	)
}

func Error(responseWriter http.ResponseWriter, statusCode int, message string) {
	Response(
		responseWriter,
		statusCode,
		StatusResponse{
			Status:  "error",
			Message: message,
		},
	)
}

func Response(responseWriter http.ResponseWriter, statusCode int, reponse StatusResponse) {
	JSON(responseWriter, statusCode, reponse)
}

func DecodeJSON(request *http.Request, data interface{}) error {
	return json.NewDecoder(request.Body).Decode(data)
}
