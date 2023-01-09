package main

import (
	"C"
	"fmt"
)

// go build -buildmode=c-archive main.go

//export ShowData
func ShowData(data int) {
	fmt.Printf("The value of x %d\n", data)
}

func main() {

}
