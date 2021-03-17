package middlewares

import (
	"fmt"
	"net/http"
)

// Logger function to log client request method and url to server
func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v \t %v \n", r.Method, r.URL.String())
		h.ServeHTTP(w, r)
	})
}