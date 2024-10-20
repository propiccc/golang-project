package view

import (
	"bufio"
	"fmt"
	"itats-project/models"
	"itats-project/services"
	"os"
)

func ShowMember(members []models.Member) {
	fmt.Println("Daftar Member : ")
	for index, item := range members {
		fmt.Printf("No: %d, ID: %d, Name: %s, Email: %s\n", index+1, item.ID, item.Name, item.Email)
	}
}

func StoreMember() models.Member {
	var id int
	var name, email string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan ID: ")
	fmt.Scan(&id)

	scanner.Scan()
	fmt.Print("Masukkan Nama: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Masukkan Email: ")
	if scanner.Scan() {
		email = scanner.Text()
	}
	return models.Member{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func EditMember() services.EditDataMember {
	var id, index int
	var name, email string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pilih List Member Yang Ingin Di Edit : ")
	fmt.Scan(&index)

	fmt.Print("Masukkan ID Terbaru: ")
	fmt.Scan(&id)

	scanner.Scan()
	fmt.Print("Masukkan Nama Terbaru: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Masukkan Email Terbaru: ")
	if scanner.Scan() {
		email = scanner.Text()
	}

	MemberData := models.Member{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return services.EditDataMember{
		Member: MemberData,
		Index:  index - 1,
	}
}

func DeleteMember() int {
	var index int

	fmt.Println("----------------------------------------------------")
	fmt.Print("Masukan List Member Yang Ingin Di Hapus : ")
	fmt.Scan(&index)
	fmt.Println("----------------------------------------------------")

	return index
}
