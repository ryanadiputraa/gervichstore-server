package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ryanadiputraa/gervichstore-server/config"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"github.com/ryanadiputraa/gervichstore-server/models"
)

type productsApi interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	UpdateAllProduct(w http.ResponseWriter, r *http.Request)
}

// ProductHandlers handle api for /api/products
type ProductHandlers struct {}

// NewProductHandlers is a constructor function for ProductHandlers
func NewProductHandlers() *ProductHandlers {
	return &ProductHandlers{}
}

// GetAllProducts is an api handler to serve all products in db
func (*ProductHandlers) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	var products models.Products

	for rows.Next() {
		var product models.Product
		rows.Scan(&product.ID, &product.Image, &product.Name, &product.Price, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt)
		products = append(products, product)
	}

	jsonBytes, err := json.Marshal(products)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	defer rows.Close()
	defer db.Close()

	return
}

// CreateProduct is an api handler to post new product to db
func(*ProductHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	newProduct := models.NewProduct()
	err = json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "bad request")
		return
	}
	
	query := `INSERT INTO products (image, name, price, stock, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Exec(query, newProduct.Image, newProduct.Name, newProduct.Price, newProduct.Stock, newProduct.Category, newProduct.CreatedAt, newProduct.UpdatedAt)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")	
		return
	}

	response := helpers.SuccessMessageFormat {
		Code: http.StatusCreated,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return	
	}
	
	helpers.WriteResponse(w, r, "application/json", http.StatusCreated, jsonBytes)
	defer db.Close()

	return
}

// GetProduct is an api handler to find certain product in db based on product id
func(*ProductHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}

	productId := strings.Split(r.URL.String(), "/")[3]
	print(productId)
	// row, err := db.Query("SELECT * FROM products WHERE id=%v", productId)
	
	return
}

// UpdateProduct is an api handler to update certain product in db based on product id
func(*ProductHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	return

}

// DeleteProduct is an api handler to delete certain product in db based on product id
func(*ProductHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	return

}