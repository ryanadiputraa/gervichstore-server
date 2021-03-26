package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ItemsController is a controller for /api/items route
func ItemsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.GetAllItems(w, r)
	default:
		helpers.WriteErrorResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}