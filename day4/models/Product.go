package models

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
	//exportable field
	ProductId int32
	Name      string
	Cost      int64
	//composition
	DOP Date
}

type Products []Product

//total no of elements
func (products Products) Len() int {
	return len(products)
}

//swap
func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//attribue/property
func (products Products) Less(i, j int) bool {
	return products[i].Name < products[j].Name
}
