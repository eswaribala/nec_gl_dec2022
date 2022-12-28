package main

import (
	"fmt"
	"unicode"
)

// Main function
func main() {

	// Creating a slice of rune
	val := []rune{'g', 'E', '3', 'K', '1'}

	// Checking each element of the given
	// slice of the rune is a letter or not
	// Using IsLetter() function
	for i := 0; i < len(val); i++ {

		if unicode.IsLetter(val[i]) == true {

			fmt.Println("It is a letter")

		} else {

			fmt.Println("It is not a letter")
		}
	}
}
