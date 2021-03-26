package helpers

import (
	"encoding/json"
	"net/http"
)

// WriteResponse is a helper for writing http response
func WriteResponse(w http.ResponseWriter, r *http.Request, contentType string, statusCode int, jsonBytes []byte) {
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}

// WriteInternalServerError is a helper for writing internal server error message
func WriteInternalServerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}

func WriteErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, errorMessage string) {
	response := ErrorMessageFormat {
		Code: statusCode,
		Error: errorMessage,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		WriteInternalServerError(w, r)
		return
	}

	WriteResponse(w, r, "application/json", statusCode, jsonBytes)
}