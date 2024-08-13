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
	s.conn.AutoMigrate(models.Customer{})

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
		return nil, repositories.ErrUsernameAlreadyTake
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
func (s *MysqlStore) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	err := s.conn.Create(customer).Error
	if err != nil {
		return nil, repositories.ErrCustomerCreate
	}
	return customer, nil
}

func (s *MysqlStore) UpdateCustomer(customerId uint, customer *models.Customer) (*models.Customer, error) {
	err := s.conn.Where("id=?", customerId).Updates(customer).Error
	if err != nil {
		return nil, repositories.ErrCustomerUpdate
	}
	return customer, nil
}
func (s *MysqlStore) GetCusotmer(id uint) (*models.Customer, error) {
	customer := &models.Customer{}
	err := s.conn.First(customer, id).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (s *MysqlStore) DeleteCusotmer(id uint) error {
	return s.conn.Model(&models.Customer{}).Delete("id=?", id).Error
}
