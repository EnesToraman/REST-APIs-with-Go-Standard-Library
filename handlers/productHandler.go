package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"main/data"
	"main/models"
)

// ProductHandler handles all the requests are made with containing "/products/" as an URL path.
// Allowed HTTP methods are GET, POST, PUT, DELETE.
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetProducts(w, r)
	case http.MethodPost:
		CreateProduct(w, r)
	case http.MethodPut:
		UpdateProduct(w, r)
	case http.MethodDelete:
		DeleteProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetProducts gets all the product information from Products array and prints them. Function is called via HTTP GET method.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Products)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "id":
			GetProductByID(w, r)
		case "name":
			GetProductByName(w, r)
		case "price":
			GetProductsByPrice(w, r)
		case "stock":
			GetProductsByStock(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// CreateProduct creates a product with the given information. Function is called via HTTP POST method.
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	data.Products = append(data.Products, product)
}

// UpdateProduct updates the product information with the given ID. Function is called via HTTP PUT method.
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	for i := range data.Products {
		if data.Products[i].ID == product.ID {
			data.Products[i].SKU = product.SKU
			data.Products[i].Name = product.Name
			data.Products[i].Price = product.Price
			data.Products[i].Stock = product.Stock
		}
	}
}

// DeleteProduct deletes the product information with the given ID. Function is called via HTTP DELETE method.
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	index := -1
	for i := range data.Products {
		if data.Products[i].ID == product.ID {
			index = i
		}
	}
	data.Products = append(data.Products[:index], data.Products[index+1:]...)
}

// GetProductByID gets the product by ID field provided in URL query.
// URL should be built with the query "?id=" and the corresponding value in order to get the related product.
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	index := -1
	if id != "" {
		productID, err := strconv.Atoi(id)
		CheckError(err)
		for i := range data.Products {
			if data.Products[i].ID == productID {
				index = i
			}
		}
		res, err := json.Marshal(data.Products[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetProductByName gets the product by Name field provided in URL query.
// URL should be built with the query "?name=" and the corresponding value in order to get the related product.
func GetProductByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	index := -1
	if name != "" {
		for i := range data.Products {
			if data.Products[i].Name == name {
				index = i
			}
		}
		res, err := json.Marshal(data.Products[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetProductsByPrice gets the product(s) by Price field provided in URL query.
// URL should be built with the query "?price=" and the corresponding value in order to get the related product(s).
func GetProductsByPrice(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	price := r.FormValue("price")
	if price != "" {
		productPrice, err := strconv.ParseFloat(price, 64)
		CheckError(err)
		for i := range data.Products {
			if data.Products[i].Price == productPrice {
				products = append(products, data.Products[i])
			}
		}
		res, err := json.Marshal(products)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetProductsByStock gets the product(s) by Stock field provided in URL query.
// URL should be built with the query "?stock=" and the corresponding value in order to get the related product(s).
func GetProductsByStock(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	stock := r.FormValue("stock")
	if stock != "" {
		productstock, err := strconv.Atoi(stock)
		CheckError(err)
		for i := range data.Products {
			if data.Products[i].Stock == productstock {
				products = append(products, data.Products[i])
			}
		}
		res, err := json.Marshal(products)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}
