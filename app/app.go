package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}