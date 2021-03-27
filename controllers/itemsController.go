package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ItemsController is a controller for /api/items route
func ItemsController(w http.ResponseWriter, r *http.Request) {
	itemHandlers := api.NewItemHandlers()

	switch r.Method {
	case "GET":
		itemHandlers.GetAllItems(w, r)
	default:
		helpers.WriteErrorResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}