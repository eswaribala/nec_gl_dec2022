package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	data := make([]int, 10)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100000)
	}
	//print before sort
	for _, value := range data {
		fmt.Printf("%d\t", value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	//print after sort
	fmt.Println("\n")
	for _, value := range data {
		fmt.Printf("%d\t", value)
	}
}
