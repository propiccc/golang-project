package main

import (
	"fmt"
	"itats-project/controllers"
)

func main() {
	var menu int
	var sys bool = true

	for sys {
		fmt.Println("-----------------------------")
		fmt.Println("1. List Member")
		fmt.Println("2. Create Member")
		fmt.Println("3. Update Member")
		fmt.Println("4. Delete Member")
		fmt.Println("-----------------------------")
		fmt.Println("5. List Product")
		fmt.Println("6. Create Product")
		fmt.Println("7. Update Product")
		fmt.Println("8. Delete Product")
		fmt.Println("10. Exit")

		menu = 100
		fmt.Println("-----------------------------")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)
		fmt.Println("-----------------------------")

		switch menu {
		case 1:
			controllers.ListMember()
		case 2:
			controllers.CreateMember()
		case 3:
			controllers.UpdateMember()
		case 4:
			controllers.DeleteMember()
		case 5:
			controllers.ListProduct()
		case 6:
			controllers.CreateProduct()
		case 7:
			controllers.UpdateProduct()
		case 8:
			controllers.DeleteProduct()
		case 10:
			sys = false
		default:
			sys = false
			fmt.Println("Menu Tidak Ada")
		}

	}

}
