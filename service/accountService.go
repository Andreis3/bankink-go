package service

import (
	"time"

	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/dto"
	"github.com/santos/banking-go/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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

func (d DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := d.repo.FindBy(req.AccountID)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := d.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := dto.MapToResponse(transaction)
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo:            repo,
		accountResponse: dto.NewAccountResponse{},
	}
}
