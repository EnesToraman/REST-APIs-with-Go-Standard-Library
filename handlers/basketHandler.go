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

// BasketHandler handles all the requests are made with containing "/baskets/" as an URL path.
// Allowed HTTP methods are GET, POST, PUT, DELETE.
func BasketHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetBaskets(w, r)
	case http.MethodPost:
		CreateBasket(w, r)
	case http.MethodPut:
		UpdateBasket(w, r)
	case http.MethodDelete:
		DeleteBasket(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetBaskets gets all the basket information from Baskets array and prints them. Function is called via HTTP GET method.
func GetBaskets(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Baskets)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "userID":
			GetBasketByUserID(w, r)
		case "productID":
			GetBasketsByProductID(w, r)
		case "sku":
			GetBasketsBySKU(w, r)
		case "quantity":
			GetBasketsByQuantity(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// CreateBasket creates a basket with the given information. Function is called via HTTP POST method.
func CreateBasket(w http.ResponseWriter, r *http.Request) {
	var basket models.Basket
	err := json.NewDecoder(r.Body).Decode(&basket)
	CheckError(err)
	data.Baskets = append(data.Baskets, basket)
}

// UpdateBasket updates the basket information with the given ID. Function is called via HTTP PUT method.
func UpdateBasket(w http.ResponseWriter, r *http.Request) {
	var basket models.Basket
	err := json.NewDecoder(r.Body).Decode(&basket)
	CheckError(err)
	for i := range data.Baskets {
		if data.Baskets[i].UserID == basket.UserID {
			data.Baskets[i].Products = basket.Products
		}
	}
}

// DeleteBasket deletes the basket information with the given ID. Function is called via HTTP DELETE method.
func DeleteBasket(w http.ResponseWriter, r *http.Request) {
	var basket models.Basket
	err := json.NewDecoder(r.Body).Decode(&basket)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	index := -1
	for i := range data.Baskets {
		if data.Baskets[i].UserID == basket.UserID {
			index = i
		}
	}
	data.Baskets = append(data.Baskets[:index], data.Baskets[index+1:]...)
}

// GetBasketByUserID gets the basket by UserID field provided in URL query.
// URL should be built with the query "?userID=" and the corresponding value in order to get the related basket.
func GetBasketByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	index := -1
	if userID != "" {
		id, err := strconv.Atoi(userID)
		CheckError(err)
		for i := range data.Baskets {
			if data.Baskets[i].UserID == id {
				index = i
			}
		}
		res, err := json.Marshal(data.Baskets[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetBasketsByProductID gets the basket(s) by ProductID field provided in URL query.
// URL should be built with the query "?productID=" and the corresponding value in order to get the related basket(s).
func GetBasketsByProductID(w http.ResponseWriter, r *http.Request) {
	var baskets []models.Basket
	productID := r.FormValue("productID")
	if productID != "" {
		basketProductID, err := strconv.Atoi(productID)
		CheckError(err)
		for i := range data.Baskets {
			for j := range data.Baskets[i].Products {
				if data.Baskets[i].Products[j].ProductID == basketProductID {
					baskets = append(baskets, data.Baskets[i])
				}
			}
		}
		res, err := json.Marshal(baskets)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetBasketsBySKU gets the basket(s) by SKU field provided in URL query.
// URL should be built with the query "?sku=" and the corresponding value in order to get the related basket(s).
func GetBasketsBySKU(w http.ResponseWriter, r *http.Request) {
	var baskets []models.Basket
	sku := r.FormValue("sku")
	if sku != "" {
		for i := range data.Baskets {
			for j := range data.Baskets[i].Products {
				if data.Baskets[i].Products[j].SKU == sku {
					baskets = append(baskets, data.Baskets[i])
				}
			}
		}
		res, err := json.Marshal(baskets)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetBasketsByQuantity gets the basket(s) by Quantity field provided in URL query.
// URL should be built with the query "?quantity=" and the corresponding value in order to get the related basket(s).
func GetBasketsByQuantity(w http.ResponseWriter, r *http.Request) {
	var baskets []models.Basket
	quantity := r.FormValue("quantity")
	if quantity != "" {
		basketQuantity, err := strconv.Atoi(quantity)
		CheckError(err)
		for i := range data.Baskets {
			for j := range data.Baskets[i].Products {
				if data.Baskets[i].Products[j].Quantity == basketQuantity {
					baskets = append(baskets, data.Baskets[i])
				}
			}
		}
		res, err := json.Marshal(baskets)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}
