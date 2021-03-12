package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}