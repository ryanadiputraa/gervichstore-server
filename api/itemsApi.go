package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/models"
)

// GetAllItems is an api handler to serve all items in db
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := []models.Items {}
	items = append(items, models.Items{})

	response := helpers.ItemsResponseFormat {
		Code: http.StatusOK,
		Data: items,
	}
	
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return	
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	return
}

// CreateItem is an api handler to post new item to db
func CreateItem(w http.ResponseWriter, r *http.Request) {
	// newItem := models.Items {}
	return

}

// GetItem is an api handler to find certain item in db based on item id
func GetItem(w http.ResponseWriter, r *http.Request) {
	return

}

// UpdateItem is an api handler to update certain item in db based on item id
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	return

}

// DeleteItem is an api handler to delete certain item in db based on item id
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	return

}