package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ryanadiputraa/gervichstore-server/helpers"
)

// FileServeHandler is a function to serve file with given filename
func FileServeHandler(w http.ResponseWriter, r *http.Request) {
	filename := helpers.GetURLParams(r, 1)

	img, err := os.Open(fmt.Sprintf("./uploads/%v", filename))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusNotFound, "no image found")
		return
	}
	
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
	defer img.Close()
	return
}