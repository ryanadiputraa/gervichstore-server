package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/controllers"
	"github.com/ryanadiputraa/gervichstore-server/middlewares"
)

// THINGS TO DO
// - manage image from server for products api

func main() {

	// routes
	http.HandleFunc("/api/products", middlewares.Logger(controllers.ProductsController))
	http.HandleFunc("/api/products/", middlewares.Logger(controllers.ProductController))
	
	// files serve
	http.HandleFunc("/", api.FileServeHandler)

	// 404
	http.HandleFunc("/api", middlewares.Logger(controllers.NotFound))
	http.HandleFunc("/api/", middlewares.Logger(controllers.NotFound))

	log.Fatal(http.ListenAndServe(":8080", nil))
}