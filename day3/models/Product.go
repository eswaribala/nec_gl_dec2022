package models

import (
	"fmt"
	"necdec2022/day3/dao"
	"strconv"
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
	Create()
	//View abstract method
	View(permission bool)
}

func (product Product) Create() {
	date := strconv.Itoa(int(product.DOP.Day)) + "/" + strconv.Itoa(int(product.DOP.Month)) + "/" + strconv.Itoa(int(product.DOP.Year))
	result, err := dao.CreateProduct(product.ProductId, product.Name, product.Cost, date)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

// View implementation methods
//call by value
func (product Product) View(permission bool) {
	if permission {
		dao.ViewProducts()
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
