package domain

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
