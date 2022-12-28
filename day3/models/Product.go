package models

import "fmt"

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
	View()
}

//implementation methods

func (product *Product) View() {
	fmt.Printf("Product%+v", product)
}
func (category *Category) View() {
	fmt.Printf("Category%+v", category)
}

//virtual inheritance pending????
