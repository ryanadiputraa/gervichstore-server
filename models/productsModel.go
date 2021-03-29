package models

import "time"

type Product struct {
	ID string `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Products []Product

func NewProduct() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}