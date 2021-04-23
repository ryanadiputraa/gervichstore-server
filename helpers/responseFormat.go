package helpers

import "github.com/ryanadiputraa/gervichstore-server/models"

// SuccessMessageFormat is a http response format for success status
type SuccessMessageFormat struct {
	Code int `json:"code"`
}

// ErrorMessageFormat is a http response format for error status
type ErrorMessageFormat struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

// ProductsResponseFormat is a http response format for array of product
type ProductsResponseFormat struct {
	Code int `json:"code"`
	Data models.Products `json:"data"`
}

// ProductResponseFormat is a http response format for single product
type ProductResponseFormat struct {
	Code int `json:"code"`
	Data models.Product `json:"data"`
}
