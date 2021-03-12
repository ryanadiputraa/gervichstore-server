package helpers

import "github.com/ryanadiputraa/gervichstore-server/models"

type ItemsResponseFormat struct {
	Code int `json:"code"`
	Data []models.Items `json:"data"`
}

type ItemResponseFormat struct {
	Code int `json:"code"`
	Data models.Items `json:"data"`
}