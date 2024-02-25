package utils

import (
	"encoding/json"
	"http-server/src/responses"
	"net/http"
)

func SetResponse(w http.ResponseWriter, httpStatus int, message string, data interface{}) {
	w.WriteHeader(httpStatus)
	response := responses.UserResponses{Status: httpStatus, Message: message, Data: map[string]interface{}{"value": data}}
	json.NewEncoder(w).Encode(response)
}
