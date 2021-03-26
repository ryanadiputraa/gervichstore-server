package helpers

import "github.com/ryanadiputraa/gervichstore-server/models"

// ErrorMessageFormat is a http response format for error status
type ErrorMessageFormat struct {
	Code int `json:"code"`
	Error string `json:"message"`
}

// ItemsResponseFormat is a http response format for array of item
type ItemsResponseFormat struct {
	Code int `json:"code"`
	Data []models.Item `json:"data"`
}

// ItemResponseFormat is a http response format for single item
type ItemResponseFormat struct {
	Code int `json:"code"`
	Data models.Item `json:"data"`
}
