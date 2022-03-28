package service

import (
	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/dto"
	"github.com/santos/banking-go/errs"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/santos/banking-go/service CustomerService
type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo             domain.CustomerRepository
	customerResponse dto.CustomerResponse
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	customerResponse := &dto.CustomerResponse{}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, customerResponse.MapCustomerToCustomerResponse(&c))
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := s.customerResponse.MapCustomerToCustomerResponse(c)

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo:             repository,
		customerResponse: dto.CustomerResponse{},
	}
}
