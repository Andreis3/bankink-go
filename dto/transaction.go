package dto

import (
	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerID      string  `json:"customer_id"`
}

type TransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t TransactionRequest) IsTransactionTypeDeposit() bool {
	return t.TransactionType == DEPOSIT
}

func (t TransactionRequest) Validate() *errs.AppError {
	if !t.IsTransactionTypeWithdrawal() && !t.IsTransactionTypeDeposit() {
		return errs.NewValidateError("Transaction type can only be deposit or withdrawal")
	}

	if t.Amount < 0 {
		return errs.NewValidateError("Amount cannot be less than zero")
	}
	return nil
}

func MapToResponse(t *domain.Transaction) TransactionResponse {
	return TransactionResponse{
		TransactionID:   t.TransactionID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
