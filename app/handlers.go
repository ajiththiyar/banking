package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ajiththiyar/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := ch.service.GetAllCustomers()
	if err != nil {
		panic(err)
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(c)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
		}
	}
}
