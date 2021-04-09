package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	// Define header what type we are going to use - in our case json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
