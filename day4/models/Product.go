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
