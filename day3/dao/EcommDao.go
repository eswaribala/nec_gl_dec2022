package dao

import (
	"database/sql"
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
