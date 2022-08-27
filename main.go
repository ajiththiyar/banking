package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullname"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {
	// define routes
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/customers", customers)

	// listen on port
	http.ListenAndServe("localhost:8080", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func customers(w http.ResponseWriter, r *http.Request) {
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
