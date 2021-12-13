package main

import (
	"net/http"

	"main/handlers"
)

func main() {
	http.HandleFunc("/users/", handlers.UserHandler)
	http.HandleFunc("/products/", handlers.ProductHandler)
	http.HandleFunc("/payments/", handlers.PaymentHandler)
	http.HandleFunc("/categories/", handlers.CategoryHandler)
	http.HandleFunc("/customers/", handlers.CustomerHandler)
	http.HandleFunc("/brands/", handlers.BrandHandler)
	http.HandleFunc("/baskets/", handlers.BasketHandler)

	err := http.ListenAndServe(":8080", nil)
	handlers.CheckError(err)
}
