package models

type Product struct {
	Name     string
	ID       int
	Quantity int
	Harga    int
}

var Products = []Product{
	{ID: 121, Name: "Ice Tea", Quantity: 10, Harga: 10000},
	{ID: 122, Name: "Lemon Tea", Quantity: 10, Harga: 10000},
	{ID: 123, Name: "Moctail", Quantity: 10, Harga: 10000},
}
