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

// PaymentHandler handles all the requests are made with containing "/payments/" as an URL path.
// Allowed HTTP methods are GET, POST, PUT, DELETE.
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetPayments(w, r)
	case http.MethodPost:
		CreatePayment(w, r)
	case http.MethodPut:
		UpdatePayment(w, r)
	case http.MethodDelete:
		DeletePayment(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetPayments gets all the payment information from Payments array and prints them. Function is called via HTTP GET method.
func GetPayments(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Payments)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "userID":
			GetPaymentsByUserID(w, r)
		case "amount":
			GetPaymentsByAmount(w, r)
		case "discount":
			GetPaymentsByDiscount(w, r)
		case "tax":
			GetPaymentsByTax(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// CreatePayment creates a payment with the given information. Function is called via HTTP POST method.
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	CheckError(err)
	data.Payments = append(data.Payments, payment)
}

// UpdatePayment updates the payment information with the given ID. Function is called via HTTP PUT method.
func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	CheckError(err)
	for i := range data.Payments {
		if data.Payments[i].ID == payment.ID {
			data.Payments[i].UserID = payment.UserID
			data.Payments[i].Amount = payment.Amount
			data.Payments[i].Discount = payment.Discount
			data.Payments[i].Tax = payment.Tax
		}
	}
}

// DeletePayment deletes the payment information with the given ID. Function is called via HTTP DELETE method.
func DeletePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	index := -1
	for i := range data.Payments {
		if data.Payments[i].ID == payment.ID {
			index = i
		}
	}
	data.Payments = append(data.Payments[:index], data.Payments[index+1:]...)
}

// GetPaymentsByUserID gets the payment(s) by UserID field provided in URL query.
// URL should be built with the query "?userID=" and the corresponding value in order to get the related payment(s).
func GetPaymentsByUserID(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	userID := r.FormValue("userID")
	if userID != "" {
		id, err := strconv.Atoi(userID)
		CheckError(err)
		for i := range data.Payments {
			if data.Payments[i].UserID == id {
				payments = append(payments, data.Payments[i])
			}
		}
		res, err := json.Marshal(payments)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetPaymentsByAmount gets the payment(s) by Amount field provided in URL query.
// URL should be built with the query "?amount=" and the corresponding value in order to get the related payment(s).
func GetPaymentsByAmount(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	amount := r.FormValue("amount")
	if amount != "" {
		paymentAmount, err := strconv.ParseFloat(amount, 64)
		CheckError(err)
		for i := range data.Payments {
			if data.Payments[i].Amount == paymentAmount {
				payments = append(payments, data.Payments[i])
			}
		}
		res, err := json.Marshal(payments)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetPaymentsByDiscount gets the payment(s) by Discount field provided in URL query.
// URL should be built with the query "?discount=" and the corresponding value in order to get the related payment(s).
func GetPaymentsByDiscount(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	discount := r.FormValue("discount")
	if discount != "" {
		paymentDiscount, err := strconv.ParseFloat(discount, 64)
		CheckError(err)
		for i := range data.Payments {
			if data.Payments[i].Discount == paymentDiscount {
				payments = append(payments, data.Payments[i])
			}
		}
		res, err := json.Marshal(payments)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetPaymentsByTax gets the payment(s) by Tax field provided in URL query.
// URL should be built with the query "?tax=" and the corresponding value in order to get the related payment(s).
func GetPaymentsByTax(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	tax := r.FormValue("tax")
	if tax != "" {
		paymentTax, err := strconv.ParseFloat(tax, 64)
		CheckError(err)
		for i := range data.Payments {
			if data.Payments[i].Tax == paymentTax {
				payments = append(payments, data.Payments[i])
			}
		}
		res, err := json.Marshal(payments)
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}
