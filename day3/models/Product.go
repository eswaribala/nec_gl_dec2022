package models

import (
	"fmt"
)

type Date struct {
	Day   int8
	Month int8
	Year  int16
}

type Category struct {
	CategoryId int32
	Name       string
}

type Product struct {
	ProductId int32
	Name      string
	Cost      int64
	//composition
	DOP Date
}

// IEcommFacade interface
type IECommFacade interface {

	//View abstract method
	View(permission bool)
}

// View implementation methods
//call by value
func (product Product) View(permission bool) {
	if permission {
		fmt.Printf("Product%+v", product)
	} else {
		fmt.Println("Not Authenticated....")
	}

}

// View call by address
func (category *Category) View(permission bool) {
	if permission {
		fmt.Printf("Category%+v", category)
	} else {
		fmt.Println("Not Authenticated....")
	}
}

//virtual inheritance pending????
