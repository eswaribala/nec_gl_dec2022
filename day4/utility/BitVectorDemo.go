package main

import (
	"fmt"
	"github.com/teivah/bitvector"
)

func main() {

	var bv1 bitvector.Ascii
	var bv2 bitvector.Ascii
	bv1 = bv1.Set(2, true)
	fmt.Println(bv1)
	bv2 = bv2.Set(1, true)
	fmt.Println(bv2)
	//masking

	bv3 := bv1.And(bv2)
	fmt.Println(bv3)

}
