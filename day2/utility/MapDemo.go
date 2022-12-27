package main

import (
	"fmt"
	"math/rand"
	"necdec2022/day2/models"
)

func main() {
	//define map
	users := make(map[int]models.User)

	//5 users collection

	for i := 0; i < 5; i++ {
		users[i] = models.User{UserId: rand.Int31n(100000), FirstName: "Parameswari",
			LastName: "Bala", MiddleName: "",
			DOB: models.Date{Day: int8(rand.Int31n(26) + 1),
				Month: int8(rand.Int31n(10) + 1), Year: 1900 + int16(rand.Int31n(122))}}
	}

	keys := make([]int, 5)
	for pos, value := range users {
		keys[pos] = pos
		fmt.Printf("%+v", value)
	}
	//search
	value, ok := users[keys[rand.Int31n(5)]]
	fmt.Println("\nKey present or not", ok)
	fmt.Printf("value%+v", value)

}
