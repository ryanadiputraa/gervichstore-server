package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
	"github.com/ryanadiputraa/gervichstore-server/middlewares"
)

// THINGS TO DO
// - manage image from server for products api

func main() {

	// routes
	http.HandleFunc("/api/products", middlewares.Logger(controllers.ProductsController))
	http.HandleFunc("/api/products/", middlewares.Logger(controllers.ProductController))

	// 404
	http.HandleFunc("/", middlewares.Logger(controllers.NotFound))

	log.Fatal(http.ListenAndServe(":8080", nil))
}