package main

import (
	"fmt"
	"log"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Product is being created....")
}
func EditProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product is being edited....")
}
func DisplayProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Retrieving Products....")
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product is being deleted....")
}

const (
	HOST_NAME = "localhost"
	HOST_PORT = "6062"
)

func main() {
	http.HandleFunc("/create", CreateProduct)
	http.HandleFunc("/update", EditProduct)
	http.HandleFunc("/display", DisplayProduct)
	http.HandleFunc("/delete", DeleteProduct)

	err := http.ListenAndServe(HOST_NAME+":"+HOST_PORT, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
