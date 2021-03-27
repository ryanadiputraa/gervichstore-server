package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/config"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/models"
)

type itemsApi interface {
	CreateItem(w http.ResponseWriter, r *http.Request)
	GetAllItems(w http.ResponseWriter, r *http.Request)
	GetItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	UpdateAllItem(w http.ResponseWriter, r *http.Request)
}

// ItemHandlers handle api for /api/items
type ItemHandlers struct {}

// NewItemHandlers is a constructor function for ItemHandlers
func NewItemHandlers() *ItemHandlers {
	return &ItemHandlers{}
}

// GetAllItems is an api handler to serve all items in db
func (*ItemHandlers) GetAllItems(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	var items models.Items

	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.Image, &item.Name, &item.Price, &item.Stock, &item.Category, &item.CreatedAt, &item.UpdatedAt)
		items = append(items, item)
	}

	jsonBytes, err := json.Marshal(items)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	defer rows.Close()
	defer db.Close()

	return
}

// CreateItem is an api handler to post new item to db
func(*ItemHandlers) CreateItem(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	newItem := models.NewItem()
	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "bad request")
		return
	}
	
	query := `INSERT INTO items (id, image, name, price, stock, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = db.Exec(query, newItem.ID, newItem.Image, newItem.Name, newItem.Price, newItem.Stock, newItem.Category, newItem.CreatedAt, newItem.UpdatedAt)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
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
func(*ItemHandlers) GetItem(w http.ResponseWriter, r *http.Request) {
	return

}

// UpdateItem is an api handler to update certain item in db based on item id
func(*ItemHandlers) UpdateItem(w http.ResponseWriter, r *http.Request) {
	return

}

// DeleteItem is an api handler to delete certain item in db based on item id
func(*ItemHandlers) DeleteItem(w http.ResponseWriter, r *http.Request) {
	return

}