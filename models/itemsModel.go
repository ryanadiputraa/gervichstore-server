package models

type Items struct {
	ID string `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	Category string `json:"category"`
}