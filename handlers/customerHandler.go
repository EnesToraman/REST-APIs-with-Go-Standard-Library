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

// CustomerHandler handles all the requests are made with containing "/customers/" as an URL path.
// Allowed HTTP methods are GET, POST, PUT, DELETE.
func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetCustomers(w, r)
	case http.MethodPost:
		CreateCustomer(w, r)
	case http.MethodPut:
		UpdateCustomer(w, r)
	case http.MethodDelete:
		DeleteCustomer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetCustomers gets all the customer information from Customers array and prints them. Function is called via HTTP GET method.
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Customers)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "id":
			GetCustomerByID(w, r)
		case "userID":
			GetCustomerByUserID(w, r)
		case "purchaseAmount":
			GetCustomersByPurchaseAmount(w, r)
		case "orderQty":
			GetCustomersByOrderQty(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// CreateCustomer creates a customer with the given information. Function is called via HTTP POST method.
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	CheckError(err)
	data.Customers = append(data.Customers, customer)
}

// UpdateCustomer updates the customer information with the given ID. Function is called via HTTP PUT method.
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	CheckError(err)
	for i := range data.Customers {
		if data.Customers[i].ID == customer.ID {
			data.Customers[i].UserID = customer.UserID
			data.Customers[i].OrderQty = customer.OrderQty
			data.Customers[i].PurchaseAmount = customer.PurchaseAmount
		}
	}
}

// DeleteCustomer deletes the customer information with the given ID. Function is called via HTTP DELETE method.
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	index := -1
	for i := range data.Customers {
		if data.Customers[i].ID == customer.ID {
			index = i
		}
	}
	data.Customers = append(data.Customers[:index], data.Customers[index+1:]...)
}

// GetCustomerByID gets the customer by ID field provided in URL query.
// URL should be built with the query "?id=" and the corresponding value in order to get the related customer.
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	index := -1
	if id != "" {
		userID, err := strconv.Atoi(id)
		CheckError(err)
		for i := range data.Customers {
			if data.Customers[i].ID == userID {
				index = i
			}
		}
		res, err := json.Marshal(data.Customers[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetCustomerByUserID gets the customer by UserID field provided in URL query.
// URL should be built with the query "?userID=" and the corresponding value in order to get the related customer.
func GetCustomerByUserID(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	userID := r.FormValue("userID")
	if userID != "" {
		id, err := strconv.Atoi(userID)
		CheckError(err)
		for i := range data.Customers {
			if data.Customers[i].UserID == id {
				customers = append(customers, data.Customers[i])
			}
		}
		res, err := json.Marshal(customers)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetCustomersByPurchaseAmount gets the customer(s) by PurchaseAmount field provided in URL query.
// URL should be built with the query "?purchaseAmount=" and the corresponding value in order to get the related customer(s).
func GetCustomersByPurchaseAmount(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	purchaseAmount := r.FormValue("purchaseAmount")
	if purchaseAmount != "" {
		customerPurchaseAmount, err := strconv.ParseFloat(purchaseAmount, 64)
		CheckError(err)
		for i := range data.Customers {
			if data.Customers[i].PurchaseAmount == customerPurchaseAmount {
				customers = append(customers, data.Customers[i])
			}
		}
		res, err := json.Marshal(customers)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetCustomersByOrderQty gets the customer(s) by Stock field provided in URL query.
// URL should be built with the query "?orderQty=" and the corresponding value in order to get the related customer(s).
func GetCustomersByOrderQty(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	orderQty := r.FormValue("orderQty")
	if orderQty != "" {
		customerOrderQty, err := strconv.Atoi(orderQty)
		CheckError(err)
		for i := range data.Customers {
			if data.Customers[i].OrderQty == customerOrderQty {
				customers = append(customers, data.Customers[i])
			}
		}
		res, err := json.Marshal(customers)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}
