package dao

import (
	"database/sql"
	"fmt"
	"necdec2022/day3/dto"

	//blank import -> init()
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DBHelper() *sql.DB {
	db, err := sql.Open("mysql",
		"root:vignesh@(127.0.0.1:3306)/ecommercedb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
func CreateProduct(ProductId int32, Name string, Cost int64, DOP string) (int64, error) {

	db := DBHelper()
	defer db.Close()
	queryString := "Insert into Product (ProductId,Name,Cost,DOP) values(?,?,?,?)"
	result, err := db.Exec(queryString, ProductId, Name, Cost, DOP)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	return result.RowsAffected()

}
func ViewProducts() {

	db := DBHelper()
	defer db.Close()
	queryString := "Select * from Product;"
	var productResponse dto.Product
	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	} else {
		for rows.Next() {
			rows.Scan(&productResponse.ProductId,
				&productResponse.Name, &productResponse.Cost, &productResponse.DOP)
			fmt.Printf("%v\n", productResponse)
		}

	}

}
