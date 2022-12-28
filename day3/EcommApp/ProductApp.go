package main

import (
	"fmt"
	"necdec2022/day3/models"
	"unsafe"
)

func main() {

	product := new(models.Product)
	product.ProductId = 845843
	product.Name = "Laptop"
	product.Cost = 3429765
	product.DOP = models.Date{Day: 12, Month: 10, Year: 2022}
	fmt.Printf("Product=%+v", product)
	fmt.Println(unsafe.Sizeof(product))

	//testing
	/*
		var user models.User
		user = models.User{UserId: 3486586, FirstName: "Parameswari",
			LastName: "Bala", MiddleName: "",
			DOB: models.Date{Day: 2, Month: 12, Year: 1970}}
		fmt.Println(unsafe.Sizeof(user))
	*/
}
