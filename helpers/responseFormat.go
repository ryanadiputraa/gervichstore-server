package helpers

import "github.com/ryanadiputraa/gervichstore-server/models"

// ItemsResponseFormat is a http response format for array of item
type ItemsResponseFormat struct {
	Code int `json:"code"`
	Data []models.Items `json:"data"`
}

// ItemResponseFormat is a http response format for single item
type ItemResponseFormat struct {
	Code int `json:"code"`
	Data models.Items `json:"data"`
}