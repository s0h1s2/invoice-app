package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type Store interface {
	FindUserByUsername(username string) (*models.User, error)
}
