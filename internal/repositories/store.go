package repositories

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
)

var (
	ErrUsernameAlreadyTake = errors.New("username already taken")
	ErrCustomerCreate      = errors.New("can't create customer")
	ErrCustomerUpdate      = errors.New("can't update customer")
)

type Store interface {
	/// User
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(username, password string) (*models.User, error)

	/// Customer
	CreateCustomer(customer models.Customer) (*models.Customer, error)
	UpdateCustomer(customerId uint, customer models.Customer) (*models.Customer, error)
	GetCusotmer(id uint) (*models.Customer, error)
	DeleteCusotmer(id uint) error
}
