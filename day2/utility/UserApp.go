package main

import (
	"fmt"
	"necdec2022/day2/models"
)

func main() {
	//create user instance
	var user models.User
	user = models.User{3486586, "Parameswari",
		"Bala", "",
		models.Date{2, 12, 1970}}
	fmt.Printf("User Data%v\n", user)

}
