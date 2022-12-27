package main

import (
	"fmt"
	"math/rand"
)

func generateOTP(seed int32) func() int32 {

	return func() int32 {
		return rand.Int31n(seed)
	}
}

func main() {

	result := generateOTP(300)
	fmt.Println(result())
	result = generateOTP(5000)
	fmt.Println(result())
	result = generateOTP(1000)
	fmt.Println(result())
}
