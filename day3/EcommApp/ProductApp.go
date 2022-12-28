package main

import (
	"fmt"
	"necdec2022/day3/models"
)

func main() {

	product := new(models.Product)
	product.ProductId = 845843
	product.Name = "Laptop"
	product.Cost = 3429765
	product.DOP = models.Date{Day: 12, Month: 10, Year: 2022}
	fmt.Printf("Product=%+v", product)
}
