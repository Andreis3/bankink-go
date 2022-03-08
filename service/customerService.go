package service

import (
	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/dto"
	"github.com/santos/banking-go/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo             domain.CustomerRepository
	customerResponse dto.CustomerResponse
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := s.customerResponse.MapCustomerToCustomerResponse(c)

	return response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo:             repository,
		customerResponse: dto.CustomerResponse{},
	}
}
