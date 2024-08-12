package repositories

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
)

var (
	UsernameAlreadyTakeErr = errors.New("Username already taken")
)

type Store interface {
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(username, password string) (*models.User, error)
}
