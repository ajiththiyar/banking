package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajiththiyar/banking/domain"
	"github.com/ajiththiyar/banking/service"
	"github.com/gorilla/mux"
)

func RunApp() {
	router := mux.NewRouter()
	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/welcome", Welcome).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customer/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)

	// listen on port
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post request received")
}
