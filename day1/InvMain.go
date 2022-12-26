package main

import (
	"fmt"
	"math/rand"
)

/*
Parameswari Ettiappan
26/12/2022
*/

//Init will not take any input parameter and not return anything
func init() {
	fmt.Println("IP initialization")

}
func init() {
	fmt.Println("Port initialization")
}
func init() {
	fmt.Println("DB Server initialization")
}

//global variable

func main() {
	var (
		productId int32
		name      string
		cost      float64
	)
	productId = rand.Int31n(1000)
	fmt.Println("Enter Name")
	fmt.Scanln(&name)

	//Type conversion
	cost = float64(rand.Int31())

	//fmt.Println("Enter Cost")
	//fmt.Scanln(&cost)
	//name = "Laptop"
	//cost = 87000
	fmt.Printf("Product Id = %d\n", productId)
	fmt.Printf("Name=%s\n", name)
	fmt.Printf("Cost=%f\n", cost)

	fmt.Println("Rocking with go!!!!")
}
