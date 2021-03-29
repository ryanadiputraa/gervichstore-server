package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/middlewares"
)

// THINGS TO DO
// - handle items api get, getAll, update, delete
// - get image from server for GetItems api
// - store image to server for CreateItem api

func main() {

	// routes
	http.HandleFunc("/api/products", middlewares.Logger(controllers.ProductsController))
	http.HandleFunc("/api/products/", middlewares.Logger(controllers.ProductController))

	// 404
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := helpers.ErrorMessageFormat {
			Code: 404,
			Error: "not found",
		}
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			helpers.WriteInternalServerError(w, r)
			return
		}
		helpers.WriteResponse(w, r, "application/json", http.StatusNotFound, jsonBytes)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}