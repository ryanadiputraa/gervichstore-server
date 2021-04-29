package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/models"
)

// NotFound is a controller for 404 routes
func NotFound(w http.ResponseWriter, r *http.Request) {
	response := models.ErrorMessageFormat {
		Code: 404,
		Error: "not found",
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}
	helpers.WriteResponse(w, r, "application/json", http.StatusNotFound, jsonBytes)
}