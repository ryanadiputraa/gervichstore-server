package main

import (
	"log"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
	"github.com/ryanadiputraa/gervichstore-server/controllers"
	"github.com/ryanadiputraa/gervichstore-server/middlewares"
)

func main() {
	// routes
	http.HandleFunc("/api/login", middlewares.Logger(controllers.Oauth2Controller))
	http.HandleFunc("/api/login/callback", middlewares.Logger(controllers.LoginCallbackController))
	http.HandleFunc("/api/products", middlewares.Logger(controllers.ProductsController))
	http.HandleFunc("/api/products/", middlewares.Logger(controllers.ProductController))
	
	// files serve
	http.HandleFunc("/", api.FileServeHandler)
	
	// 404
	http.HandleFunc("/api", middlewares.Logger(controllers.NotFound))
	http.HandleFunc("/api/", middlewares.Logger(controllers.NotFound))

	log.Fatal(http.ListenAndServe(":8080", nil))
}