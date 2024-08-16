// package mysql
package mysqlstore

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
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
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
func (s *userStore) CreateUser(username, password string) (*models.User, error) {
	newUser := models.User{
		Username: username,
		Password: password,
	}
	err := s.conn.db.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (s *userStore) UpdateUserPassword(id uint, password string) (*models.User, error) {
	return nil, nil
}
