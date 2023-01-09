package main

//#cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "customer.h"
/*

extern int displayCustomer(struct Customer *customer, char *out);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	name := C.CString("Parameswari")
	defer C.free(unsafe.Pointer(name))

	customerId := C.int(382548)
	day := C.int(2)
	month := C.int(12)
	year := C.int(1970)

	d := C.struct_DOB{
		day:   day,
		month: month,
		year:  year,
	}

	customerData := C.struct_Customer{
		name:       name,
		customerId: customerId,
		dob:        d,
	}

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.displayCustomer(&customerData, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}
