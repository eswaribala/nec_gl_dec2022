package main

import (
	"github.com/gorilla/mux"

	_ "necdec2022/day7/docs"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"necdec2022/day7/stores"
	"net/http"
)

// @title Customer API
// @version 1.0
// @description This is api service for managing Customer
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email parameswaribala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:6064
// @BasePath /
func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/customers", stores.CreateCustomer).Methods("POST")
	// Read
	router.HandleFunc("/customers/{accountNo}", stores.GetCustomerById).Methods("GET")
	// Read-all
	router.HandleFunc("/customers", stores.GetCustomers).Methods("GET")
	// Update
	router.HandleFunc("/customers", stores.UpdateCustomer).Methods("PUT")
	// Delete
	router.HandleFunc("/customers/{accountNo}", stores.DeleteCustomerById).Methods("DELETE")
	// Initialize db connection
	stores.InitDB()

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//router.PathPrefix("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(http.ListenAndServe(":6064", router))
}
