package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ItemController is a controller for /api/item route
func ItemController(w http.ResponseWriter, r *http.Request) {
	itemHandlers := api.NewItemHandlers()

	switch r.Method {
	case "GET":
		itemHandlers.GetItem(w, r)	
	case "POST":
		itemHandlers.CreateItem(w, r)
	case "PUT":
		itemHandlers.UpdateItem(w, r)
	case "DELETE":
		itemHandlers.DeleteItem(w, r)
	default:
		helpers.WriteErrorResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}