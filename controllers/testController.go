package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
) 

func TestController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.GetTest(w, r)
	}
}