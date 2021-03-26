package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/config"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/models"
)

// GetAllItems is an api handler to serve all items in db
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := models.Items{}
	newItem := models.NewItem()
	items = append(items, *newItem)

	response := helpers.ItemsResponseFormat {
		Code: http.StatusOK,
		Data: items,
	}
	
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return	
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	return
}

// CreateItem is an api handler to post new item to db
func CreateItem(w http.ResponseWriter, r *http.Request) {
	db := config.OpenConnection()

	newItem := models.NewItem()
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "bad request")
		return
	}
	
	query := `INSERT INTO items (id, image, name, price, stock, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = db.Exec(query, newItem.ID, newItem.Image, newItem.Name, newItem.Price, newItem.Stock, newItem.Category, newItem.CreatedAt, newItem.UpdatedAt)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusInternalServerError, "error")
		return
	}

	response := helpers.SuccessMessageFormat {
		Code: http.StatusCreated,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return	
	}
	
	helpers.WriteResponse(w, r, "application/json", http.StatusCreated, jsonBytes)
	defer db.Close()

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