package main

import "fmt"

func main() {

	countries := make([]string, 4, 6)
	//confirm append beyond capacity???

	//existing allocation, we cannot increase capacity
	countries = countries[4:10]
	countries[3] = "Singapore"
	//new allocation, we can increase capacity and data will be lost
	countries = make([]string, 4, 10)
	fmt.Println(countries[3])
	fmt.Println(cap(countries))
	fmt.Println(len(countries))
}
