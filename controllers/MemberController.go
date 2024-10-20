package controllers

import (
	"fmt"
	view "itats-project/View"
	"itats-project/services"
)

func ListMember() {
	view.ShowMember(services.GetMember())
}

func CreateMember() {
	newMember := view.StoreMember()
	services.CreateMember(newMember)
	ListMember()
}

func UpdateMember() {
	ListMember()
	fmt.Println("----------------------------------------------------")
	updateMember := view.EditMember()
	services.UpdateMember(updateMember)
	fmt.Println("----------------------------------------------------")
	ListMember()
	fmt.Println("----------------------------------------------------")

}

func DeleteMember() {
	ListMember()
	index := view.DeleteMember()
	services.DestroyMember(index)
}
