package dto

import "github.com/santos/banking-go/domain"

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

func (cr CustomerResponse) statusAsText() string {
	statusText := "active"
	if cr.Status == "0" {
		statusText = "inactive"
	}

	return statusText
}

func (cr CustomerResponse) MapCustomerToCustomerResponse(c *domain.Customer) CustomerResponse {

	return CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      cr.statusAsText(),
	}
}
