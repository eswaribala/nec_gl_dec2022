package main

import (
	"math/rand"
	"necdec2022/day2/models"
)

var users map[int]models.User

func init() {
	//define map
	users = make(map[int]models.User)

	//5 users collection

	for i := 0; i < 5; i++ {
		users[i] = models.User{UserId: int32(i), FirstName: "Parameswari",
			LastName: "Bala", MiddleName: "",
			DOB: models.Date{Day: int8(rand.Int31n(26) + 1),
				Month: int8(rand.Int31n(10) + 1), Year: 1900 + int16(rand.Int31n(122))}}
	}

}

func checkUserId(userId int32) bool {

	if _, ok := users[int(userId)]; !ok {
		return false
	}
	return true
}
