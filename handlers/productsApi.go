package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	}
	defer db.Close()

	var query string
	productNameQuery, ok := r.URL.Query()["product_name"]
	if ok {
		query = fmt.Sprintf("SELECT * FROM products WHERE LOWER(name) LIKE LOWER('%v%%')", productNameQuery[0])
	} else {
		query = "SELECT * FROM products"
	}

	rows, err := db.Query(query)
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
	if products == nil {
		products = models.Products{}
	}

	response := models.ProductsResponseFormat {
		Code: http.StatusOK,
		Data: products,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	defer rows.Close()
	return
}

// GetProduct is an api handler to find certain product in db based on product id
func(*ProductHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
	}
	defer db.Close()

	productId := helpers.GetURLParams(r, 3)
	row, err := db.Query(fmt.Sprintf("SELECT * FROM products WHERE id=%v", productId))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")	
		return	
	}

	var product models.Product
	isNotNull := row.Next()
	if !isNotNull {
		helpers.WriteErrorResponse(w, r, http.StatusNotFound, "not found")	
		return	
	}
	row.Scan(&product.ID, &product.Image, &product.Name, &product.Price, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt)

	response := models.ProductResponseFormat {
		Code: http.StatusOK,
		Data: product,
	}
	
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return	
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	defer row.Close()
	return
}

// CreateProduct is an api handler to post new product to db
func(*ProductHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
	}
	defer db.Close()

	// assign product data
	newProduct := models.NewProduct()
	newProduct.Name = r.FormValue("name")
	newProduct.Price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "price must be int")
		return
	}
	newProduct.Stock, err = strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "stock must be int")
		return
	}
	newProduct.Category = r.FormValue("category")
	
	// store and assign product image
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "bad request")
		return
	}
	defer file.Close()

	newProduct.Image = fmt.Sprintf("%v/%v", os.Getenv("BASE_URL"), handler.Filename)
	fileLocation := filepath.Join("uploads", handler.Filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}

	// insert product to database
	query := `INSERT INTO products (name, image, price, stock, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Exec(query, newProduct.Name, newProduct.Image, newProduct.Price, newProduct.Stock, newProduct.Category, newProduct.CreatedAt, newProduct.UpdatedAt)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "bad request")	
		return
	}

	response := models.SuccessMessageFormat {
		Code: http.StatusCreated,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return	
	}
	
	helpers.WriteResponse(w, r, "application/json", http.StatusCreated, jsonBytes)
	return
}

// UpdateProduct is an api handler to update certain product in db based on product id
func(*ProductHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
	}
	defer db.Close()

	// delete current image
	productId := helpers.GetURLParams(r, 3)
	row, err := db.Query(fmt.Sprintf("SELECT image FROM products WHERE id = %v", productId))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "provide an id")	
		return	
	}
	var currentProduct models.Product
	isNotNull := row.Next()
	if !isNotNull {
		helpers.WriteErrorResponse(w, r, http.StatusNotFound, "not found")	
		return	
	}
	row.Scan(&currentProduct.Image)
	currentProduct.Image = strings.Split(currentProduct.Image, "/")[1]
	if _, err := os.Stat(fmt.Sprintf("%v/uploads/%v",os.Getenv("WORKING_DIRECTORY") ,currentProduct.Image)); !os.IsNotExist(err) {
		fmt.Print("File exist")
		if err = os.Remove(fmt.Sprintf("%v/uploads/%v",os.Getenv("WORKING_DIRECTORY"), currentProduct.Image)); err != nil {
			helpers.WriteInternalServerError(w, r)
			return
		}
	}
	
	// assign product data
	var updatedProduct models.Product
	updatedProduct.Name = r.FormValue("name")
	updatedProduct.Price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "price must be int")
		return
	}
	updatedProduct.Stock, err = strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "stock must be int")
		return
	}
	updatedProduct.Category = r.FormValue("category")
	
	// store and assign product image
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadRequest, "bad request")
		return
	}
	defer file.Close()

	updatedProduct.Image = fmt.Sprintf("%v/%v", os.Getenv("BASE_URL"), handler.Filename)
	fileLocation := filepath.Join("uploads", handler.Filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}

	query := fmt.Sprintf("UPDATE products SET image = '%v', name = '%v', price = %v, stock = %v, category = '%v', updated_at = '%v' WHERE id = %v", updatedProduct.Image, updatedProduct.Name, updatedProduct.Price, updatedProduct.Stock, updatedProduct.Category, time.Now().Format(time.RFC3339), productId)
	_, err = db.Exec(query)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")	
		return
	}

	response := models.SuccessMessageFormat {
		Code: http.StatusCreated,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return	
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusAccepted, jsonBytes)
	return
}

// DeleteProduct is an api handler to delete certain product in db based on product id
func(*ProductHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenConnection()
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
	}
	defer db.Close()
	
	productId := helpers.GetURLParams(r, 3)

	// check if product exist
	row, err := db.Query(fmt.Sprintf("SELECT image FROM products WHERE id = %v", productId))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")	
		return	
	}
	isNotNull := row.Next()
	if !isNotNull {
		helpers.WriteErrorResponse(w, r, http.StatusNotFound, "not found")	
		return	
	}

	// Delete product that exist
	_, err = db.Query(fmt.Sprintf("DELETE FROM products WHERE id = %v", productId))
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "bad gateway")
		return
	}
	var productImage string
	row.Scan(&productImage)
	productImage = strings.Split(productImage, "/")[1]
	if _, err := os.Stat(fmt.Sprintf("%v/uploads/%v",os.Getenv("WORKING_DIRECTORY") ,productImage)); !os.IsNotExist(err) {
		fmt.Print("File exist")
		if err = os.Remove(fmt.Sprintf("%v/uploads/%v",os.Getenv("WORKING_DIRECTORY"), productImage)); err != nil {
			helpers.WriteInternalServerError(w, r)
			return
		}
	}

	response := models.SuccessMessageFormat {
		Code: http.StatusOK,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		helpers.WriteInternalServerError(w, r)
		return
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, jsonBytes)
	return
}