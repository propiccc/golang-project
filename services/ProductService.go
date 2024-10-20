package services

import (
	"fmt"
	"itats-project/models"
)

type EditDataProduct struct {
	Product models.Product
	Index   int
}

func GetProduct() []models.Product {
	return models.Products
}

func CreateProduct(newData models.Product) {
	models.Products = append(models.Products, newData)
	fmt.Println("Member Created")
}

func UpdateProduct(updateData EditDataProduct) {
	if updateData.Index >= 0 && updateData.Index < len(models.Products) {
		models.Products[updateData.Index] = updateData.Product
		fmt.Println("Member berhasil diperbarui!")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func DestroyProduct(index int) {
	if index > len(models.Products) || index < 0 {
		fmt.Println("Data Member Tidak Di Temukan!")
		return
	}

	var UpdateData []models.Product
	for i, item := range models.Products {
		if i != index-1 {
			UpdateData = append(UpdateData, item)
		}
	}

	models.Products = UpdateData
	fmt.Println("Berhasil Di Delete!")
}
