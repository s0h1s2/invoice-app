package repositories

import (
	"errors"
)

var (
	ErrUsernameAlreadyTake = errors.New("username already taken")
	ErrCustomerCreate      = errors.New("can't create customer")
	ErrCustomerUpdate      = errors.New("can't update customer")
	ErrNotFound            = errors.New("not found")
)
