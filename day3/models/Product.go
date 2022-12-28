package models

type Date struct {
	Day   int8
	Month int8
	Year  int16
}

type Product struct {
	ProductId int32
	Name      string
	Cost      int64
	DOP       Date
}
