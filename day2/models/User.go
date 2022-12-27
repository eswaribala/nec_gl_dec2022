package models

//sharable attributes should always follow
//every words first letter in upper case

type User struct {
	UserId     int32
	FirstName  string
	LastName   string
	MiddleName string
	DOB        Date
}
