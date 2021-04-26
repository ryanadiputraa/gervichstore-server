package middlewares

import (
	"fmt"
	"net/http"

	"gopkg.in/gookit/color.v1"
)

// Logger function to log client request method and url to server
func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set background color based on http status code response
		// statusCode := w.Header().Get("status")
		statusCode := 200
		bgColor := color.HEX("976024", true) // default bgColor
		
		if statusCode >= 200 && statusCode < 300 {
			bgColor = color.HEX("#0c771a", true)
		} else if statusCode >= 500 && statusCode < 600 {
			bgColor = color.HEX("ac1b1b", true)
		}
		resStatusCode := bgColor

		// log message to console
		color.Info.Printf("%v \t", r.Method)
		fmt.Printf("%v \t", r.URL.String())
		resStatusCode.Println(statusCode)
		h.ServeHTTP(w, r)
	})
}