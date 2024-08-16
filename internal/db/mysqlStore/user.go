// package mysql
package mysqlstore

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"gorm.io/gorm"
)

type userStore struct {
	conn *mysqlStore
}

func NewMysqlUserStore(conn *mysqlStore) *userStore {
	return &userStore{
		conn: conn,
	}
}

func (s *userStore) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.conn.db.First(&user, "username=?", username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrInvalidCreds
		}
		return nil, err
	}
	return &user, nil
}
func (s *userStore) CreateUser(newUser *models.User) (*models.User, error) {
	err := s.conn.db.Model(&models.User{}).Create(newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *userStore) UpdateUserPassword(user *models.User) (*models.User, error) {
	return nil, nil
}
func (s *userStore) CreateSession(session *models.Session) error {
	err := s.conn.db.Create(session).Error
	return err
}

func (s *userStore) GetSession(token string) (*models.Session, error) {
	return nil, nil
}
