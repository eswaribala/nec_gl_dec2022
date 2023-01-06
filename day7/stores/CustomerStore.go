package stores

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"necdec2022/day7/models"
)

var db *gorm.DB

func InitDB() {
	var err error
	dataSourceName := "root:vignesh@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	//db, err = gorm.Open("mysql", "root:vignesh@(localhost:3306)/nec_trader_db?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.Exec("Create Database NECCustomerDB2023")
	db.Exec("use NECCustomerDB2023")
	//generate the tables
	db.AutoMigrate(&models.Customer{})

}

// CreateCustomer godoc
// @Summary Create a new Customer
// @Description Create a new customer with the input payload
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.Customer true "Create Customer"
// @Success 200 {object} models.Customer
// @Router /customers [post]

func CreateCustomer() {

}

// put
func UpdateCustomer() {

}

// get
func GetCustomers() {

}

// get
func GetCustomerById() {

}

// delete
func DeleteCustomerById() {

}
