package helpers

import (
	"net/http"
	"strconv"

	"gopkg.in/gookit/color.v1"
)

func logResponseStatusCode(w http.ResponseWriter, r *http.Request) {
	// set background color based on http status code response
	// statusCode := r.Header
	statusCode, err := strconv.Atoi(w.Header().Get("status"))
	if err != nil {
		statusCode = 0
	}
	bgColor := color.HEX("976024", true) // default bgColor
	
	if statusCode >= 200 && statusCode < 300 {
		bgColor = color.HEX("#0c771a", true)
	} else if statusCode >= 500 && statusCode < 600 {
		bgColor = color.HEX("ac1b1b", true)
	}
	resStatusCode := bgColor
	resStatusCode.Println(statusCode)
}