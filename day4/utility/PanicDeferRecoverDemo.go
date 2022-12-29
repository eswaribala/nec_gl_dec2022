package main

import "fmt"

var (
	accountNo int32
)

func finally() {
	accountNo = 0
}

func RecoverAccount() {

	//catch block
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func ValidateAccount(accountNo int32, password string) {
	defer RecoverAccount()
	//finally
	defer finally()
	if accountNo < 0 {
		panic("Account No should be 18 digit number") //try block
	}

	if len(password) == 0 {
		panic("Password cannot be null") //try block
	}

	fmt.Printf("%d,%s\n", accountNo, password)
	fmt.Println("Validate function return Normal")

}

func main() {
	accountNo = -8468746
	defer fmt.Println("Running at the end from main")
	ValidateAccount(accountNo, "35696")
	//fmt.Println("Running at the end from main")
}
