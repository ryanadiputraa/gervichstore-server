package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
)

// THINGS TO DO
// - handle itemsApi for itemController

func main() {

	// routes
	http.HandleFunc("/api/items", controllers.ItemsController)
	http.HandleFunc("/api/items/", controllers.ItemController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}