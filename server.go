package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/middlewares"
)

// THINGS TO DO
// - handle items api get, getAll, update, delete

func main() {

	// routes
	http.HandleFunc("/api/items", middlewares.Logger(controllers.ItemsController))
	http.HandleFunc("/api/items/", middlewares.Logger(controllers.ItemController))

	// 404
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.WriteResponse(w, r, "application/json", http.StatusNotFound, []byte("not found"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}