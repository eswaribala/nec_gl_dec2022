package main

import (
	"html/template"
	"log"
	"math/rand"
	"necdec2022/day4/models"
	"net/http"
	"strconv"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "7070"
)

//function
func renderTemplate(w http.ResponseWriter, r *http.Request) {
	product := models.Product{
		ProductId: rand.Int31n(100000),
		Name:      "Product" + strconv.Itoa(int(rand.Int31n(100000))),
		DOP: models.Date{
			Day:   1 + int8(rand.Int31n(26)),
			Month: 1 + int8(rand.Int31n(11)),
			Year:  1900 + int16(rand.Int31n(122)),
		},
		Cost: int64(rand.Int31n(100000)),
	}
	parsedTemplate, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplate.Execute(w, product)
	if err != nil {
		log.Printf("Error occurred while executing the templateor writing its output : ", err)
		return
	}
}
func main() {
	http.HandleFunc("/home", renderTemplate)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
