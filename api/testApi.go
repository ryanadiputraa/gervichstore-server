package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/models"
)

// GetTest handle test request
func GetTest(w http.ResponseWriter, r *http.Request) {
	person := models.Person {
		Name: "name",
	}
	response, err := json.Marshal(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return	
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}