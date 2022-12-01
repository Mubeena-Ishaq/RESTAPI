package methods

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var latestCusId int = 0

// Create Customer
func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db, err := connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Creating Customer Instance
	var customer *Customer
	// customer := Customer{
	// 	CustomerId: ,
	// }
	//{
	// 	for _, Customer := range customer {
	// 		customer.customerId
	// 	  }
	// }

	// var customers = &Customer{}
	// for _, customers := range customers.First_Name {
	// 	customers.CustomerId()
	// }

	// Decoding Request
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	latestCusId++
	customer.Id = latestCusId

	// Inserting Into Database
	_, err = db.Model(customer).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Returning Customer
	json.NewEncoder(w).Encode(customer)
}

// Get Customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db, err := connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()
	//Creating Customers Slice
	var customers []Customer
	if err := db.Model(&customers).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Returning Customers
	json.NewEncoder(w).Encode(customers)
}

// Get Customer
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db, err := connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	customerId := params["id"]
	ci, _ := strconv.Atoi(customerId)

	// u, _ := url.Parse(r.URL.String())
	// q := u.Query()
	// customerId := q["id"]
	// //params.Get()
	// ci, _ := strconv.Atoi(customerId)

	// Creating Customer Instance
	customer := &Customer{Id: ci}
	if err := db.Model(customer).WherePK().Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	//Returning Customer
	json.NewEncoder(w).Encode(customer)
}

// Update Customer
func allupdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db, err := connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	customerId := params["id"]
	ci, _ := strconv.Atoi(customerId)

	// Creating Customer Instance
	customer := &Customer{Id: ci}

	_ = json.NewDecoder(r.Body).Decode(&customer)
	if customer.FirstName == "" {
		log.Println("Can't update")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if customer.LastName == "" {
		log.Println("Can't update")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if customer.Gender == "" {
		log.Println("Can't update")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if customer.Age == 0 {
		log.Println("Can't update")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.Model(customer).WherePK().Update(&customer)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Returning Customer
	json.NewEncoder(w).Encode(customer)
}

// Update Customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db, err := connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	customerId := params["id"]
	ci, _ := strconv.Atoi(customerId)

	// Creating Customer Instance
	customer := &Customer{Id: ci}

	_ = json.NewDecoder(r.Body).Decode(&customer)
	_, err = db.Model(customer).WherePK().Set("first_name = ?, last_name = ?, age = ?, gender = ?", customer.FirstName, customer.LastName, customer.Age, customer.Gender).Update()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Returning Customer
	json.NewEncoder(w).Encode(customer)
}

// Delete Customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db, err := connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	customerId := params["id"]
	ci, _ := strconv.Atoi(customerId)
	// Creating Customer Instance
	// customer := &Customer{Id: ci}
	// result, err := db.Model(customer).WherePK().Delete()

	// Creating Customer Instance Alternative Way
	customer := &Customer{}
	result, err := db.Model(customer).Where("id=?", ci).Delete()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	//Returning Customer
	json.NewEncoder(w).Encode(result)
}
