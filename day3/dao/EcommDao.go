package dao

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"necdec2022/day3/dto"
	"os"

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

	//image
	/*
		fileToBeUploaded := "nec.png"
		file, err := os.Open(fileToBeUploaded)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close()

		fileInfo, _ := file.Stat()
		var size int64 = fileInfo.Size()
		bytes1 := make([]byte, size)

		// read file into bytes
		buffer := bufio.NewReader(file)
		result1, err := buffer.Read(bytes1) // <--------------- here!

		fmt.Println(result1)
	*/
	// Create blob
	fileToBeUploaded := "logo.png"
	r, e := os.Open(fileToBeUploaded)
	if e != nil {
		panic(e)
	}
	defer r.Close()
	var dest []byte
	blob := bytes.NewBuffer(dest)
	io.Copy(blob, r)

	db := DBHelper()
	defer db.Close()
	queryString := "Insert into Product (ProductId,Name,Cost,DOP,Image) values(?,?,?,?,?)"
	result, err := db.Exec(queryString, ProductId, Name, Cost, DOP, blob.Bytes())
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
				&productResponse.Name, &productResponse.Cost, &productResponse.DOP, &productResponse.Image)
			fmt.Printf("%+v\n", productResponse)
		}

	}

}
