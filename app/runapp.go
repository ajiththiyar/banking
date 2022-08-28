package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunApp() {
	mux := mux.NewRouter()
	// define routes
	mux.HandleFunc("/welcome", Welcome)
	mux.HandleFunc("/customers", Customers)

	// listen on port
	log.Fatal(http.ListenAndServe("localhost:8080", mux))

}
