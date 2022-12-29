package main

import (
	"fmt"
	"necdec2022/day3/facades"
	"necdec2022/day3/models"
	"unsafe"
)

func main() {

	product := models.Product{}
	product.ProductId = 845828
	product.Name = "Mobile"
	product.Cost = 342976
	product.DOP.Day = 12
	product.DOP.Month = 11
	product.DOP.Year = 2022
	//product.Order.OrderId = 4356734
	//product.Order.Amount = 4386874
	fmt.Println(unsafe.Sizeof(product))

	var productInstance facades.IECommFacade
	productInstance = product
	//productInstance.Create()
	productInstance.View(true)

	//testing
	/*
		var user models.User
		user = models.User{UserId: 3486586, FirstName: "Parameswari",
			LastName: "Bala", MiddleName: "",
			DOB: models.Date{Day: 2, Month: 12, Year: 1970}}
		fmt.Println(unsafe.Sizeof(user))
	*/
}
