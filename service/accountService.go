package service

import (
	"time"

	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/dto"
	"github.com/santos/banking-go/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo            domain.AccountRepository
	accountResponse dto.NewAccountResponse
}

func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		OpeningDAte: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := d.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := d.accountResponse.MapAccountToAccountResponse(newAccount)

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo:            repo,
		accountResponse: dto.NewAccountResponse{},
	}
}
