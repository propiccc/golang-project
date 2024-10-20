package controllers

import (
	"fmt"
	view "itats-project/View"
	"itats-project/services"
)

func ListProduct() {
	view.ShowProduct(services.GetProduct())
}

func CreateProduct() {
	newProduct := view.StoreProduct()
	services.CreateProduct(newProduct)
	ListProduct()
}

func UpdateProduct() {
	ListProduct()
	fmt.Println("----------------------------------------------------")
	updateProduct := view.EditProduct()
	services.UpdateProduct(updateProduct)
	fmt.Println("----------------------------------------------------")
	ListProduct()
	fmt.Println("----------------------------------------------------")

}

func DeleteProduct() {
	ListProduct()
	index := view.DeleteProduct()
	services.DestroyProduct(index)
}
