package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/santos/banking-go/config"
	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/service"
)

func Start() {
	router := mux.NewRouter()

	dbClient := getDBClient()

	// wiring
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{
		service: service.NewCustomerService(customerRepositoryDb),
	}
	ah := AccountHandler{
		service: service.NewAccountService(accountRepositoryDb),
	}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT), router))
}

func getDBClient() *sqlx.DB {
	driver := config.DB_DRIVER
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	client, err := sqlx.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
