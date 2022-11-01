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
		WriteResponse(w, appError.Code, appError.EMessage(), false)
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			WriteResponse(w, http.StatusOK, c, true)
		} else {
			WriteResponse(w, http.StatusOK, c, false)
		}
	}
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	c, appError := ch.service.GetCustomer(vars["customer_id"])
	if appError != nil {
		WriteResponse(w, appError.Code, appError.EMessage(), false)
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			WriteResponse(w, http.StatusOK, c, true)
		} else {
			WriteResponse(w, http.StatusOK, c, false)
		}
	}
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}, xmlVal bool) {
	if xmlVal {
		w.Header().Add("Content-Type", "application/xml")
		w.WriteHeader(code)
		xml.NewEncoder(w).Encode(data)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(data)
	}

}
