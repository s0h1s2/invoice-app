package db

import (
	"errors"
	"log"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlStore struct {
	conn *gorm.DB
}

func NewMysqlStore(dsn string) *MysqlStore {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to DB:%s", err.Error())
	}
	return &MysqlStore{
		conn: conn,
	}

}
func (s *MysqlStore) Init() {
	s.conn.AutoMigrate(models.User{})
}
func (s *MysqlStore) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.conn.First(&user, "username=?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}
func (s *MysqlStore) CreateUser(username, password string) (*models.User, error) {
	user, _ := s.FindUserByUsername(username)
	if user != nil {
		return nil, repositories.UsernameAlreadyTakeErr
	}
	newUser := models.User{
		Username: username,
		Password: password,
	}
	err := s.conn.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
