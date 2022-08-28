package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullname" xml:"fullname"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Customers(w http.ResponseWriter, r *http.Request) {
	c := []Customer{
		{"aj", "nashik", "394585"},
		{"manish", "dhule", "365585"},
		{"anup", "mumbai", "394235"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(c)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}
}
