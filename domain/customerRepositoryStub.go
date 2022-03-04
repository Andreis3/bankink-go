package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Test_Name", "Test_City", "99999-999", "2001-01-01", "1"},
		{"1002", "Test_Name_2", "Test_City_2", "88888-888", "2001-02-02", "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
