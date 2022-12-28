package main

import (
	"fmt"
	"necdec2022/day3/models"
	"unsafe"
)

func main() {
	//it returns address
	product := new(models.Product)
	product.ProductId = 845843
	product.Name = "Laptop"
	product.Cost = 3429765
	product.DOP.Day = 12
	product.DOP.Month = 10
	product.DOP.Year = 2022

	fmt.Println(unsafe.Sizeof(product))

	var productInstance models.IECommFacade
	productInstance = product
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
