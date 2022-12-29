package main

import (
	"encoding/json"
	"fmt"
	"necdec2022/day4/models"
)

func main() {

	product := models.Product{ProductId: 237867, Name: "Door", Cost: 20000,
		DOP: models.Date{Day: 11, Month: 12, Year: 2022}}

	result, _ := json.Marshal(product)
	fmt.Printf("%s\n", result)
	/*
		var goResult map[string]interface{}

		json.Unmarshal(result, &goResult)
		fmt.Printf("%s\n", goResult)
	*/
}
