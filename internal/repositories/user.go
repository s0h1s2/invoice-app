package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type UserRepository interface {
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	CreateSession(session *models.Session) error
	GetSession(token string) (*models.Session, error)
	UpdateUserPassword(user *models.User) (*models.User, error)
}
