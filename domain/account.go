package domain

import "github.com/santos/banking-go/errs"

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDAte string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountID string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
