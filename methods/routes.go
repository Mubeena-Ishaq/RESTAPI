package methods

// import (
// 	"fmt"

// 	"github.com/google/uuid"
// )

import (
	"github.com/gorilla/mux"
	"net/http"
	//"github.com/joho/godotenv"
)

func Handler() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/customers", createCustomer).Methods(http.MethodPost)
	r.HandleFunc("/api/customers", getCustomers).Methods(http.MethodGet)
	r.HandleFunc("/api/customers/{id}", getCustomer).Methods(http.MethodGet)
	r.HandleFunc("/api/customers/{id}", deleteCustomer).Methods(http.MethodDelete)
	r.HandleFunc("/api/customers/{id}", updateCustomer).Methods(http.MethodPatch)
	r.HandleFunc("/api/customers/{id}", allupdateCustomer).Methods(http.MethodPut)
	// db := connect()
	// defer db.Close()

	// To Validate
	// customer := Customer{CustomerId: uuid.(), First_Name: "Mubeena", Last_Name: "Ishaq", Age: 25, Gender: "F"}
	// fmt.Println(customer)
	return r
}
