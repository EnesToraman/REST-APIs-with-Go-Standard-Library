package main

import (
	"net/http"
	"os"

	"main/handlers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/users/", handlers.UserHandler)
	http.HandleFunc("/products/", handlers.ProductHandler)
	http.HandleFunc("/payments/", handlers.PaymentHandler)
	http.HandleFunc("/categories/", handlers.CategoryHandler)
	http.HandleFunc("/customers/", handlers.CustomerHandler)
	http.HandleFunc("/brands/", handlers.BrandHandler)
	http.HandleFunc("/baskets/", handlers.BasketHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	handlers.CheckError(err)
}
