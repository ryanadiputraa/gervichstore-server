package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/config"
	"github.com/ryanadiputraa/gervichstore-server/handlers"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ProductsController is a controller for /api/products route
func ProductsController(w http.ResponseWriter, r *http.Request) {
	productHandlers := handlers.NewProductHandlers()
	config.SetupCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

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