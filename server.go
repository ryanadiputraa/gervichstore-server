package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
)

// THINGS TO DO
// - added db (postgresql)
// - handle itemsApi for itemController
// - added auth for create, update, and delete items

func main() {

	// routes
	http.HandleFunc("/api/items", controllers.ItemsController)
	http.HandleFunc("/api/items/", controllers.ItemController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}