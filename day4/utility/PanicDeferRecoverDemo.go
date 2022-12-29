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

func ValidatePassword(password string) {
	defer RecoverAccount()
	//finally
	defer finally()

	if len(password) == 0 {
		panic("Password cannot be null") //try block
	}

	fmt.Printf("%d,%s\n", accountNo, password)
	fmt.Println("Validate function return Normal")

}

func ValidateAccountNo(accountNo int32) {
	//defer RecoverAccount()
	//finally
	//defer finally()
	if accountNo < 0 {
		panic("Account No should be 18 digit number") //try block
	}

	fmt.Printf("%d\n", accountNo)
	fmt.Println("Validate function return Normal")

}

func main() {
	accountNo = 8468746
	defer fmt.Println("Running at the end from main")
	ValidateAccountNo(accountNo)
	ValidatePassword("")
	//fmt.Println("Running at the end from main")
}
