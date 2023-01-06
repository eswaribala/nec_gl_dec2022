package models

type Customer struct {
	AccountNo int64  `json:"accountNo" gorm:"primary_key"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	ContactNo int64  `json:"contactNo"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
