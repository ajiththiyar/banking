package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ajiththiyar/banking/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, appError := ch.service.GetAllCustomers()
	if appError != nil {
		w.WriteHeader(appError.Code)
		fmt.Fprintf(w, appError.Message)
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

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	c, appError := ch.service.GetCustomer(vars["customer_id"])
	if appError != nil {
		w.WriteHeader(appError.Code)
		fmt.Fprintf(w, appError.Message)
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
