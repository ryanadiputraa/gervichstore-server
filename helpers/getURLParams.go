package helpers

import (
	"net/http"
	"strings"
)

// GetURLParams is a helper function to get URL parameter on certain position
func GetURLParams(r *http.Request, paramsPosition int) string {
	return strings.Split(r.URL.String(), "/")[paramsPosition]
}