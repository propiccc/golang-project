package view

import (
	"bufio"
	"fmt"
	"itats-project/models"
	"itats-project/services"
	"os"
)

func ShowProduct(data []models.Product) {
	fmt.Println("Daftar Product : ")
	for index, item := range data {
		fmt.Printf("No: %d, ID: %d, Name: %s, Quantity: %d,  Harga: %d\n", index+1, item.ID, item.Name, item.Quantity, item.Harga)
	}
}

func StoreProduct() models.Product {
	var id, quantity, harga int
	var name string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan ID Product: ")
	fmt.Scan(&id)

	scanner.Scan()
	fmt.Print("Masukkan Nama Product: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Masukkan Quantity Product: ")
	fmt.Scan(&quantity)

	fmt.Print("Masukkan Harga Product: ")
	fmt.Scan(&harga)

	return models.Product{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Harga:    harga,
	}
}

func EditProduct() services.EditDataProduct {
	var id, index, harga, quantity int
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pilih List Product Yang Ingin Di Edit : ")
	fmt.Scan(&index)

	fmt.Print("Masukkan ID Terbaru: ")
	fmt.Scan(&id)

	scanner.Scan()
	fmt.Print("Masukkan Nama Terbaru: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Masukkan quantity Terbaru: ")
	fmt.Scan(&quantity)

	fmt.Print("Masukkan Harga Terbaru: ")
	fmt.Scan(&harga)

	updateData := models.Product{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Harga:    harga,
	}

	return services.EditDataProduct{
		Product: updateData,
		Index:   index - 1,
	}
}

func DeleteProduct() int {
	var index int

	fmt.Println("----------------------------------------------------")
	fmt.Print("Masukan List Member Yang Ingin Di Hapus : ")
	fmt.Scan(&index)
	fmt.Println("----------------------------------------------------")

	return index
}
