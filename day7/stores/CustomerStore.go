package stores

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
}
