package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type UserRepository interface {
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(username, password string) (*models.User, error)
	UpdateUserPassword(id uint, password string) (*models.User, error)
}
