package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/controllers"
)

func main() {
	// test routes
	http.HandleFunc("/test", controllers.TestController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}