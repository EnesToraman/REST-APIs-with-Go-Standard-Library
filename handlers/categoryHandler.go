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

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetCategories(w, r)
	case http.MethodPost:
		CreateCategory(w, r)
	case http.MethodPut:
		UpdateCategory(w, r)
	case http.MethodDelete:
		DeleteCategory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetCategories gets all the category information from Categories array and prints them. Function is called via HTTP GET method.
func GetCategories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Categories)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "id":
			GetCategoryByID(w, r)
		case "name":
			GetCategoryByName(w, r)
		case "productQty":
			GetCategoriesByProductQty(w, r)
		case "isMain":
			GetCategoriesByActivity(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// Createcategory creates a category with the given information. Function is called via HTTP POST method.
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	CheckError(err)
	data.Categories = append(data.Categories, category)
}

// UpdateCategory updates the category information with the given ID. Function is called via HTTP PUT method.
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	CheckError(err)
	for i := range data.Categories {
		if data.Categories[i].ID == category.ID {
			data.Categories[i].Name = category.Name
			data.Categories[i].ProductQty = category.ProductQty
			data.Categories[i].IsMain = category.IsMain
		}
	}
}

// DeleteCategory deletes the category information with the given ID. Function is called via HTTP DELETE method.
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	var index int
	for i := range data.Categories {
		if data.Categories[i].ID == category.ID {
			index = i
		}
	}
	data.Categories = append(data.Categories[:index], data.Categories[index+1:]...)
}

// GetCategoryByID gets the category by ID field provided in URL query.
// URL should be built with the query "?id=" and the corresponding value in order to get the related category.
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	index := -1
	if id != "" {
		categoryID, err := strconv.Atoi(id)
		CheckError(err)
		for i := range data.Categories {
			if data.Categories[i].ID == categoryID {
				index = i
			}
		}
		res, err := json.Marshal(data.Categories[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetCategoryByName gets the category by Name field provided in URL query.
// URL should be built with the query "?name=" and the corresponding value in order to get the related category.
func GetCategoryByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	index := -1
	if name != "" {
		for i := range data.Categories {
			if data.Categories[i].Name == name {
				index = i
			}
		}
		res, err := json.Marshal(data.Categories[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetCategoriesByProductQty gets the category(s) by ProductQty field provided in URL query.
// URL should be built with the query "?productQty=" and the corresponding value in order to get the related category(s).
func GetCategoriesByProductQty(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	productQty := r.FormValue("productQty")
	if productQty != "" {
		customerProductQty, err := strconv.Atoi(productQty)
		CheckError(err)
		for i := range data.Categories {
			if data.Categories[i].ProductQty == customerProductQty {
				categories = append(categories, data.Categories[i])
			}
		}
		res, err := json.Marshal(categories)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetCategoriesByMain gets the category(s) by isMain field provided in URL query.
// URL should be built with the query "?isMain=" and the corresponding value in order to get the related category(s).
func GetCategoriesByActivity(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	isMain := r.FormValue("isMain")
	main, err := strconv.ParseBool(isMain)
	CheckError(err)
	if main {
		for i := range data.Categories {
			if data.Categories[i].IsMain == main {
				categories = append(categories, data.Categories[i])
			}
		}
	} else {
		for i := range data.Categories {
			if data.Categories[i].IsMain == main {
				categories = append(categories, data.Categories[i])
			}
		}
	}
	res, err := json.Marshal(categories)
	CheckError(err)
	fmt.Fprint(w, string(res))
}
