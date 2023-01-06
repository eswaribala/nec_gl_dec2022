package stores

import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"net/http"
)

// GetEmployees godoc
// @Summary Get details of all employees
// @Description Get details of all employees
// @Tags employees
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Customer
// @Router /customers [get]
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Create a Resty Client
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("https://jsonplaceholder.typicode.com/users")
	fmt.Fprintf(w, resp.String())
}
