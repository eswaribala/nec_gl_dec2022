package main

import "necdec2022/day3/models"

func main() {
	//it returns address
	category := new(models.Category)
	category.CategoryId = 39456934
	category.Name = "Electronics"
	var categoryInstance models.IECommFacade
	categoryInstance = category
	categoryInstance.View(false)
}
