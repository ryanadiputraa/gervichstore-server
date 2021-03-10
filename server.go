package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Person is a test models
type Person struct {
	Name string `json:"name"`
}

func main() {
	// test routes
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		person := Person {
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
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}