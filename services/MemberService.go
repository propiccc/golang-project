package services

import (
	"fmt"
	"itats-project/models"
)

type EditDataMember struct {
	Member models.Member
	Index  int
}

func GetMember() []models.Member {
	return models.Members
}

func CreateMember(newMember models.Member) {
	models.Members = append(models.Members, newMember)
	fmt.Println("Member Created")
}

func UpdateMember(updateMember EditDataMember) {
	if updateMember.Index >= 0 && updateMember.Index < len(models.Members) {
		models.Members[updateMember.Index] = updateMember.Member
		fmt.Println("Member berhasil diperbarui!")
	} else {
		fmt.Println("Index tidak valid.")
	}

}

func DestroyMember(index int) {
	if index > len(models.Members) || index < 0 {
		fmt.Println("Data Member Tidak Di Temukan!")
		return
	}

	var UpdateMember []models.Member
	for i, member := range models.Members {
		if i != index-1 {
			UpdateMember = append(UpdateMember, member)
		}
	}

	models.Members = UpdateMember
	fmt.Println("Berhasil Di Delete!")
}
