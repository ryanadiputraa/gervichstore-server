package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
)

func main() {

	// routes
	http.HandleFunc("/api/items", controllers.ItemsController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}