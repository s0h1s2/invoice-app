package services

import (
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
)

type CustomerService struct {
	store repositories.Store
}

func NewCustomerService(store repositories.Store) *CustomerService {
	return &CustomerService{
		store: store,
	}
}
func (c *CustomerService) CreateCustomer(customer dto.CreateCustomerRequest) (*models.Customer, error) {
	newCustomer := models.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
		Phone:     customer.Phone,
		Balance:   customer.Balance,
	}
	result, err := c.store.CreateCustomer(&newCustomer)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *CustomerService) FindCustomerById(customerId uint) (*models.Customer, error) {
	return c.store.GetCusotmer(customerId)
}
func (c *CustomerService) UpdateCustomer(customerId uint, customer dto.UpdateCustomerRequest) (*models.Customer, error) {

	newCustomer := models.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Phone:     customer.Phone,
		Address:   customer.Address,
		Balance:   customer.Balance,
	}
	result, err := c.store.UpdateCustomer(customerId, &newCustomer)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *CustomerService) DeleteCustomer(customerId uint) error {
	return c.store.DeleteCusotmer(customerId)
}
