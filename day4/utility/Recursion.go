package main

import (
	"fmt"
	"time"
)

func OTPCountDown(number int) {

	if number > 0 {
		fmt.Println(number)
		time.Sleep(1 * time.Second)
		OTPCountDown(number - 1)
	} else {
		fmt.Println("Time Exceeded, Generate OTP")
	}

}

func main() {
	OTPCountDown(60)
}
