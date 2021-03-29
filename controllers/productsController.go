package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ProductsController is a controller for /api/products route
func ProductsController(w http.ResponseWriter, r *http.Request) {
	productHandlers := api.NewProductHandlers()

	switch r.Method {
	case "GET":
		productHandlers.GetAllProducts(w, r)
	default:
		helpers.WriteErrorResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}