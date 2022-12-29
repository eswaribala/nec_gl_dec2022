package main

import (
	"fmt"
	"github.com/teivah/bitvector"
)

func main() {

	var bv1 bitvector.Ascii
	var bv2 bitvector.Ascii
	bv1 = bv1.Set(5, true)
	bv2 = bv2.Set(6, true)

	//masking

	bv3 := bv1.Xor(bv2)
	fmt.Println(bv3)

}
