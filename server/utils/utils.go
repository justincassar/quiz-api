package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Send interfaces as JSON to client
func EncodeData(w http.ResponseWriter, data interface{}) {

	// Start encoder and encode the data interface
	err := json.NewEncoder(w).Encode(data)
	if err != nil {

		// Set header to a Server Error if encoding did not work
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// Set header data and provide Status OK
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Show bad request HTTP error
func SendBadRequest(w http.ResponseWriter) {

	// Set status header to a "Bad Request"
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	// Map response with a "Bad Request" string
	res := make(map[string]string)
	res["message"] = "Bad Request"

	// Encode error message as JSON and send to client
	json, err := json.Marshal(res)
	if err != nil {
		log.Panicf("Error in JSON marshal. Err: %s", err)
	}
	w.Write(json)
	return
}
