package data

import "main/models"

var Users []models.User = []models.User{
	{ID: 1, Email: "enestoraman@yahoo.com", Username: "ET", FirstName: "Enes", LastName: "Toraman", IsActive: true},
	{ID: 2, Email: "enstrmn@hotmail.com", Username: "EnsTrmn", FirstName: "Enes", LastName: "TRMN", IsActive: true},
	{ID: 3, Email: "toraman.enes@yahoo.com", Username: "Enes.Toraman", FirstName: "Enes", LastName: "Toraman", IsActive: false},
	{ID: 4, Email: "enestoraman@gmail.com", Username: "EnesToraman", FirstName: "Enes", LastName: "Toraman", IsActive: true},
	{ID: 5, Email: "enestoraman@hotmail.com", Username: "E.T", FirstName: "Enes", LastName: "Toraman", IsActive: false},
	{ID: 6, Email: "enes.toraman@boun.edu.tr", Username: "EnTo", FirstName: "Enes", LastName: "Toraman", IsActive: true},
}

var Products []models.Product = []models.Product{
	{ID: 1, SKU: "11004545", Name: "Samsung Television", Price: 7999.90, Stock: 67},
	{ID: 2, SKU: "22115656", Name: "iPhone 13 Pro", Price: 13999.90, Stock: 15},
	{ID: 3, SKU: "33226767", Name: "HP Laptop", Price: 10999.90, Stock: 33},
	{ID: 4, SKU: "44337878", Name: "Philips Razor", Price: 699.90, Stock: 49},
	{ID: 5, SKU: "55448989", Name: "Xiaomi Robot", Price: 4999.90, Stock: 21},
	{ID: 6, SKU: "66559090", Name: "MX Mouse", Price: 199.90, Stock: 86},
}

var Baskets []models.Basket = []models.Basket{
	{UserID: 1, Products: []models.BasketProduct{
		{ProductID: 1, SKU: "11004545", Quantity: 2},
		{ProductID: 2, SKU: "22115656", Quantity: 4},
		{ProductID: 3, SKU: "33226767", Quantity: 1},
	}},
	{UserID: 2, Products: []models.BasketProduct{
		{ProductID: 2, SKU: "22115656", Quantity: 4},
		{ProductID: 5, SKU: "55448989", Quantity: 1},
	}},
	{UserID: 5, Products: []models.BasketProduct{
		{ProductID: 1, SKU: "11004545", Quantity: 1},
		{ProductID: 2, SKU: "22115656", Quantity: 3},
		{ProductID: 3, SKU: "33226767", Quantity: 1},
		{ProductID: 6, SKU: "66559090", Quantity: 6},
	}},
}

var Payments []models.Payment = []models.Payment{
	{ID: 1, UserID: 1, Amount: 11546.50, Discount: 150.50, Tax: 48.76},
	{ID: 2, UserID: 1, Amount: 15550.80, Discount: 120.20, Tax: 156.78},
	{ID: 3, UserID: 1, Amount: 3456.78, Discount: 100.22, Tax: 97.45},
	{ID: 4, UserID: 2, Amount: 5656.12, Discount: 80.88, Tax: 67.54},
	{ID: 6, UserID: 4, Amount: 6787.35, Discount: 120.65, Tax: 14.78},
	{ID: 7, UserID: 4, Amount: 1999.90, Discount: 110.10, Tax: 66.88},
	{ID: 8, UserID: 5, Amount: 8080.80, Discount: 145.20, Tax: 84.55},
	{ID: 9, UserID: 6, Amount: 1100.56, Discount: 25.44, Tax: 99.10},
}

var Customers []models.Customer = []models.Customer{
	{ID: 1, UserID: 1, OrderQty: 3, PurchaseAmount: 30554.08},
	{ID: 2, UserID: 2, OrderQty: 1, PurchaseAmount: 5656.12},
	{ID: 3, UserID: 4, OrderQty: 2, PurchaseAmount: 12443.47},
	{ID: 4, UserID: 5, OrderQty: 1, PurchaseAmount: 8080.80},
	{ID: 5, UserID: 6, OrderQty: 1, PurchaseAmount: 1100.56},
}

var Categories []models.Category = []models.Category{
	{ID: 1, Name: "Electronic", ProductQty: 25, IsMain: true},
	{ID: 2, Name: "Laptop", ProductQty: 17, IsMain: false},
	{ID: 3, Name: "Shoe", ProductQty: 34, IsMain: true},
	{ID: 4, Name: "Self-Care", ProductQty: 5, IsMain: false},
}

var Brands []models.Brand = []models.Brand{
	{ID: 1, Name: "Samsung", ProductQty: 6, TotalWorth: 70000.15},
	{ID: 2, Name: "Apple", ProductQty: 12, TotalWorth: 55700.84},
	{ID: 3, Name: "HP", ProductQty: 15, TotalWorth: 36840.47},
	{ID: 4, Name: "Philips", ProductQty: 9, TotalWorth: 47520.32},
	{ID: 5, Name: "Xiaomi", ProductQty: 21, TotalWorth: 86000.29},
	{ID: 6, Name: "MX", ProductQty: 2, TotalWorth: 2000.24},
}
