package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
)

// ItemsController is a controller for /api/items route
func ItemsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.GetAllItems(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}