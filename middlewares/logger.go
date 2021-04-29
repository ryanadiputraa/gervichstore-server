package middlewares

import (
	"fmt"
	"net/http"

	"gopkg.in/gookit/color.v1"
)

// Logger function to log client request method and url to server
func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		

		// log message to console
		color.Info.Printf("%v \t", r.Method)
		fmt.Printf("%v \t", r.URL.String())
		// resStatusCode.Println(statusCode)
		h.ServeHTTP(w, r)
	})
}