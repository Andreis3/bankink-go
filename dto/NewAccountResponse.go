package dto

import "github.com/santos/banking-go/domain"

type NewAccountResponse struct {
	AccountID string `json:"account_id"`
}

func (n NewAccountResponse) MapAccountToAccountResponse(a *domain.Account) NewAccountResponse {
	return NewAccountResponse{
		AccountID: a.AccountID,
	}
}
