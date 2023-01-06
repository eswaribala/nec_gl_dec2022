package stores

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"necdec2022/day7/models"
	"net/http"
	"strconv"
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
	//db.Exec("Create Database NECCustomerDB2023")
	db.Exec("use NECCustomerDB2023")
	//generate the tables
	db.AutoMigrate(&models.Customer{})

}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer with the input paylod
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.Customer true "Create customer"
// @Success 200 {object} models.Customer
// @Router /customers [post]
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	// Creates new customer by inserting records in the `customers` and `items` table
	db.Create(&customer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomer godoc
// @Summary Update customer identified by the given AccountNo
// @Description Update the customer corresponding to the input AccountNo
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.Customer true "update customer"
// @Success 200 {object} models.Customer
// @Router /customers [put]
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var updatedCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	db.Save(&updatedCustomer)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCustomer)
}

// GetCustomers godoc
// @Summary Get details of all customers
// @Description Get details of all customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Customer
// @Router /customers [get]
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customer
	db.Find(&customers)
	json.NewEncoder(w).Encode(customers)
}

// GetCustomerById godoc
// @Summary Get details for a given accountNo
// @Description Get details of customer  corresponding to the input accountNo
// @Tags customers
// @Accept  json
// @Produce  json
// @Param accountNo path int true "ID of the Customer"
// @Success 200 {object} models.Customer
// @Router /customers/{accountNo} [get]
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputAccountNo := params["accountNo"]

	var customer models.Customer
	db.First(&customer, inputAccountNo)
	json.NewEncoder(w).Encode(customer)
}

// DeleteCustomerById godoc
// @Summary Delete customer identified by the given accountNo
// @Description Delete the customer corresponding to the input accountNo
// @Tags customers
// @Accept  json
// @Produce  json
// @Param accountNo path int true "ID of the customer to be deleted"
// @Success 204 "No Content"
// @Router /customers/{accountNo} [delete]
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputAccountNo := params["accountNo"]
	// Convert `traderId` string param to uint64
	id64, _ := strconv.ParseUint(inputAccountNo, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("account_no=?", idToDelete).Delete(&models.Customer{})
	w.WriteHeader(http.StatusNoContent)
}
