package domain

import "github.com/santos/banking-go/errs"

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDAte string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
