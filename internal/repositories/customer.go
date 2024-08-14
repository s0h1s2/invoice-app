package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) (*models.Customer, error)
	UpdateCustomer(customerId uint, customer *models.Customer) (*models.Customer, error)
	GetCusotmer(customerId uint) (*models.Customer, error)
	DeleteCusotmer(customerId uint) error
}
