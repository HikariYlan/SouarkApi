package send

import (
	"encoding/json"
	"net/http"
)

// Converts data into JSON and sends it with the status code
func Json(data any, res http.ResponseWriter, statusCode int) {
	// JSON data generation
	jsonData, err := json.Marshal(data)

	// Error handling
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Data type definition
	res.Header().Set("Content-Type", "application/json")

	// HTTP response code
	res.WriteHeader(statusCode)

	// Data sending
	res.Write(jsonData)
}
