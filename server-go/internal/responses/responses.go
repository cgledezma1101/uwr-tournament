package responses

import (
	"encoding/json"
	"net/http"
)

// SendJSON sends a JSON response with the specified status code
func SendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

// SendError sends an error response with the specified status code
func SendError(w http.ResponseWriter, message string, statusCode int) {
	errorResp := map[string]string{"error": message}
	SendJSON(w, errorResp, statusCode)
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Data interface{} `json:"data"`
}
