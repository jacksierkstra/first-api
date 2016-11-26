package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {

	// Always set response to JSON.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := Response{}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response.Message = "The HTTP method you've used wasn't allowed."
		response.StatusCode = http.StatusMethodNotAllowed
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "Welcome to my awesome API."
	response.StatusCode = http.StatusOK
	json.NewEncoder(w).Encode(response)

}
