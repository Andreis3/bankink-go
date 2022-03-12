package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/santos/banking-go/dto"
	"github.com/santos/banking-go/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = customerID
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

// MakeTransaction /customers/2000/accounts/90720
func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountID := vars["account_id"]
	customerID := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	// build the request object
	request.AccountID = accountID
	request.CustomerID = customerID

	// make transaction
	account, appError := h.service.MakeTransaction(request)

	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, account)
	}
}
