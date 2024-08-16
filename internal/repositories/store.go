package repositories

import (
	"errors"
)

var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrInvalidCreds         = errors.New("invalid credentials")
	ErrCustomerCreate       = errors.New("can't create customer")
	ErrCustomerUpdate       = errors.New("can't update customer")
	ErrNotFound             = errors.New("not found")
)
