package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"

	m "github.com/Mubeena-Ishaq/RESTAPI/methods"
)

func main() {

	fmt.Println(strings.EqualFold("Methods", "methods"))
	r := m.Handler()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	// Starting Server
	log.Fatal(http.ListenAndServe(":5678", r))
}
