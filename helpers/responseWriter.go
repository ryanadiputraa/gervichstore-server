package helpers

import "net/http"

// WriteResponse is a helper for writing http response
func WriteResponse(w http.ResponseWriter, r *http.Request, contentType string, statusCode int, jsonBytes []byte) {
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}