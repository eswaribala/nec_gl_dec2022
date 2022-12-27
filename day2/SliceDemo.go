package main

import "fmt"

func main() {

	countries := make([]string, 2, 6)

	fmt.Println(cap(countries))
	fmt.Println(len(countries))
}
