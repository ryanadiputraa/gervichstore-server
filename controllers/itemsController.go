package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/api"
)

func ItemsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.GetAllItems(w, r)
	}
}