package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/handlers"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ProductsController is a controller for /api/products route
func ProductsController(w http.ResponseWriter, r *http.Request) {
	productHandlers := handlers.NewProductHandlers()

	switch r.Method {
	case "GET":
		productHandlers.GetAllProducts(w, r)
	case "POST":
		productHandlers.CreateProduct(w, r)
	default:
		helpers.WriteErrorResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}