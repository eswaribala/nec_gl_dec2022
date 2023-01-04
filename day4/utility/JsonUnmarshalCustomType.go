package main

import (
	"encoding/json"
	"fmt"
	"necdec2022/day4/models"
)

func main() {
	alice := models.Category{}
	//aliceJson := `{"categoryId": 237867, "name": "Electronics"}`
	aliceJson := `{"categoryId": 237867, "name": {"commonName":"Electronics","categoryName":""}}`
	err := json.Unmarshal([]byte(aliceJson), &alice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", alice)
}
