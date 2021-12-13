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

// BrandHandler handles all the requests are made with containing "/brands/" as an URL path.
// Allowed HTTP methods are GET, POST, PUT, DELETE.
func BrandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetBrands(w, r)
	case http.MethodPost:
		CreateBrand(w, r)
	case http.MethodPut:
		UpdateBrand(w, r)
	case http.MethodDelete:
		DeleteBrand(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetBrands gets all the brand information from Brands array and prints them. Function is called via HTTP GET method.
func GetBrands(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Brands)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "id":
			GetBrandByID(w, r)
		case "name":
			GetBrandByName(w, r)
		case "productQty":
			GetBrandsByProductQty(w, r)
		case "totalWorth":
			GetBrandsByTotalWorth(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// CreateBrand creates a brand with the given information. Function is called via HTTP POST method.
func CreateBrand(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	err := json.NewDecoder(r.Body).Decode(&brand)
	CheckError(err)
	data.Brands = append(data.Brands, brand)
}

// UpdateBrand updates the brand information with the given ID. Function is called via HTTP PUT method.
func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	err := json.NewDecoder(r.Body).Decode(&brand)
	CheckError(err)
	for i := range data.Brands {
		if data.Brands[i].ID == brand.ID {
			data.Brands[i].Name = brand.Name
			data.Brands[i].ProductQty = brand.ProductQty
			data.Brands[i].TotalWorth = brand.TotalWorth
		}
	}
}

// DeleteBrand deletes the brand information with the given ID. Function is called via HTTP DELETE method.
func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	err := json.NewDecoder(r.Body).Decode(&brand)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	index := -1
	for i := range data.Brands {
		if data.Brands[i].ID == brand.ID {
			index = i
		}
	}
	data.Brands = append(data.Brands[:index], data.Brands[index+1:]...)
}

// GetBrandByID gets the brand by ID field provided in URL query.
// URL should be built with the query "?id=" and the corresponding value in order to get the related brand.
func GetBrandByID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	index := -1
	if id != "" {
		brandID, err := strconv.Atoi(id)
		CheckError(err)
		for i := range data.Brands {
			if data.Brands[i].ID == brandID {
				index = i
			}
		}
		res, err := json.Marshal(data.Brands[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetBrandByName gets the brand by Name field provided in URL query.
// URL should be built with the query "?name=" and the corresponding value in order to get the related brand.
func GetBrandByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	index := -1
	if name != "" {
		for i := range data.Brands {
			if data.Brands[i].Name == name {
				index = i
			}
		}
		res, err := json.Marshal(data.Brands[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetBrandsByProductQty gets the brand(s) by ProductQty field provided in URL query.
// URL should be built with the query "?productQty=" and the corresponding value in order to get the related brand(s).
func GetBrandsByProductQty(w http.ResponseWriter, r *http.Request) {
	var brands []models.Brand
	productQty := r.FormValue("productQty")
	if productQty != "" {
		brandProductQty, err := strconv.Atoi(productQty)
		CheckError(err)
		for i := range data.Brands {
			if data.Brands[i].ProductQty == brandProductQty {
				brands = append(brands, data.Brands[i])
			}
		}
		res, err := json.Marshal(brands)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetBrandsByTotalWorth gets the brand(s) by TotalWorth field provided in URL query.
// URL should be built with the query "?totalWorth=" and the corresponding value in order to get the related brand(s).
func GetBrandsByTotalWorth(w http.ResponseWriter, r *http.Request) {
	var brands []models.Brand
	totalWorth := r.FormValue("totalWorth")
	if totalWorth != "" {
		customerTotalWorth, err := strconv.ParseFloat(totalWorth, 64)
		CheckError(err)
		for i := range data.Brands {
			if data.Brands[i].TotalWorth == customerTotalWorth {
				brands = append(brands, data.Brands[i])
			}
		}
		res, err := json.Marshal(brands)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}
