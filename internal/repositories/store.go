package repositories

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
)

var (
	UsernameAlreadyTakeErr = errors.New("Username already taken")
	CustomerCreateErr      = errors.New("Can't create customer")
	CustomerUpdateErr      = errors.New("Can't update customer")
)

type Store interface {
	/// User
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(username, password string) (*models.User, error)

	/// Customer
	CreateCustomer(firstName, lastName, address, phone string, balance float32) (*models.Customer, error)
	UpdateCustomer(customerId uint, firstName, lastName, address, phone string, balance float32) (*models.Customer, error)
	GetCusotmer(id uint) (*models.Customer, error)
	DeleteCusotmer(id uint) error
}
