package main

import (
	"fmt"
	"necdec2022/day2/models"
)

func main() {
	//create user instance
	var user models.User
	user = models.User{UserId: 3486586, FirstName: "Parameswari",
		LastName: "Bala", MiddleName: "",
		DOB: models.Date{Day: 2, Month: 12, Year: 1970}}

	fmt.Printf("User Data%+v\n", user)

}
