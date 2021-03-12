package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/models"
)

// GetAllItems is an api to serve all items in db
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := []models.Items {}
	items = append(items, models.Items{})

	response, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return	
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, response)
	return
}