package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	fmt.Println("Testing....")
	fmt.Println(len(args))
	if len(args) == 4 {
		result, err := configureServer(args)
		if err == nil {

			//range final position excluded
			for index, value := range result[1:len(result)] {
				fmt.Println(index, value)
			}
			/*for i := 1; i < len(result); i++ {
				fmt.Println(result[i])
			}*/
		} else {
			fmt.Println(err)
		}
	}
}

func configureServer(serverProperties []string) ([]string, error) {
	//create another array
	serverProps := make([]string, 4)

	//traditional loop
	/*for i := 1; i < len(serverProperties); i++ {
		fmt.Println(serverProperties[i])
	}*/

	port, err := strconv.Atoi(serverProperties[1])
	//if there is error, return err
	if err != nil {
		return nil, err
	}
	//if value < 1027 , return err

	if port < 1027 {
		return nil, err
	} else {

		for index, value := range serverProperties {
			serverProps[index] = value
		}
		return serverProps, nil
	}
}
