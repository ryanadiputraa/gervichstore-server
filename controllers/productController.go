package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/config"
	"github.com/ryanadiputraa/gervichstore-server/handlers"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ProductController is a controller for /api/product route
func ProductController(w http.ResponseWriter, r *http.Request) {
	productHandlers := handlers.NewProductHandlers()
	config.SetupCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	switch r.Method {
	case "GET":
		// check if theres an id on URL, if not get all products
		productId := helpers.GetURLParams(r, 3)
		if len(productId) > 0 {
			productHandlers.GetProduct(w, r)
			return	
		}
		productHandlers.GetAllProducts(w, r)
	case "POST":
		productHandlers.CreateProduct(w, r)
	case "PUT":
		productHandlers.UpdateProduct(w, r)
	case "DELETE":
		productHandlers.DeleteProduct(w, r)
	default:
		helpers.WriteErrorResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}