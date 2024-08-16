package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type UserRepository interface {
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUserPassword(id uint, password string) (*models.User, error)
}
