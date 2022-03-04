package service

import "github.com/santos/banking-go/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.FindById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repository,
	}
}