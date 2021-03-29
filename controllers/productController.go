package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// ProductController is a controller for /api/product route
func ProductController(w http.ResponseWriter, r *http.Request) {
	productHandlers := api.NewProductHandlers()

	switch r.Method {
	case "GET":
		productHandlers.GetProduct(w, r)	
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