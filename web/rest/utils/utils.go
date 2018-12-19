package utils

import (
	"encoding/json"
	"net/http"
)

// Message - util for creating the body of a response
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
	}
}

// Respond - sets res to json and pipes message to response body
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
