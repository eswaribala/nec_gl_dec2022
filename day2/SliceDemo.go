package main

import (
	"fmt"
	"strconv"
)

func main() {

	countries := make([]string, 4, 6)
	//confirm append beyond capacity???

	//The append built-in function appends elements to the end of a slice. If
	// it has sufficient capacity, the destination is resliced to accommodate the
	// new elements. If it does not, a new underlying array will be allocated.
	// Append returns the updated slice. It is therefore necessary to stores the
	// result of append, often in the variable holding the slice itself:
	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)
	// As a special case, it is legal to append a string to a byte slice, like this:
	//	slice = append([]byte("hello "), "world"...)
	for i := 0; i < 10; i++ {
		//you must stores
		countries = append(countries, "country"+strconv.Itoa(i))
	}

	for _, value := range countries {
		fmt.Println(value)
	}
	//existing allocation, we cannot increase capacity
	countries = countries[4:10]
	countries[3] = "Singapore"
	countries[4] = "USA"

	for _, value := range countries {
		fmt.Println(value)
	}
	//new allocation, we can increase capacity and data will be lost
	countries = make([]string, 4, 10)
	fmt.Println(countries[3])
	fmt.Println(cap(countries))
	fmt.Println(len(countries))
}
