package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ItemController is a controller for /api/item route
func ItemController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.GetItem(w, r)	
	case "POST":
		api.CreateItem(w, r)
	case "PUT":
		api.UpdateItem(w, r)
	case "DELETE":
		api.DeleteItem(w, r)
	default:
		helpers.WriteResponse(w, r, "application/json", http.StatusMethodNotAllowed, []byte("Method not allowed"))
		return
	}
}