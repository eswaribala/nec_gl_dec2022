package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Testing....")
	fmt.Println(len(args))
	if len(args) == 4 {
		result := configureServer(args)
		for index, value := range result {
			fmt.Println(index, value)
		}
	}
}

func configureServer(serverProperties []string) []string {
	//create another array
	serverProps := make([]string, 4)
	for index, value := range serverProperties {
		serverProps[index] = value
	}
	return serverProps
}
