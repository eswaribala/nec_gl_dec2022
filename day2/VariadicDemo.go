package main

import (
	"errors"
	"fmt"
)

func main() {
	//invoking the function
	result, err := getMarketData("OPEC", 16, 17, 18, 19)
	if err == nil {
		fmt.Println(result)
	}

	getMarketData("OPEC", 16)

	result, err = getMarketData("OPEC")
	if err != nil {
		fmt.Println(err)
	}

	getMarketData("OPEC", 16, 17, 18, 19, 20)

}

//variadic function--> number of arguments
func getMarketData(opec string, price ...int32) (int32, error) {
	var sum int32 = 0
	if len(price) == 0 {
		return 0, errors.New("not having price list")
	} else {
		for _, value := range price {
			sum = sum + value
		}
	}
	return sum, nil
}
