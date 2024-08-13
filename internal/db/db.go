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
func (s *MysqlStore) CreateCustomer(firstName, lastName, address, phone string, balance float32) (*models.Customer, error) {
	customer := models.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		Phone:     phone,
		Balance:   balance,
	}
	err := s.conn.Create(&customer).Error
	if err != nil {
		return nil, repositories.CustomerCreateErr
	}
	return &customer, nil
}

func (s *MysqlStore) UpdateCustomer(customerId uint, firstName, lastName, address, phone string, balance float32) (*models.Customer, error) {
	newCustomer := models.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Balance:   balance,
		Phone:     phone,
		Address:   address,
	}
	err := s.conn.Updates(&newCustomer).Where("id=?", customerId).Error
	if err != nil {
		return nil, repositories.CustomerUpdateErr
	}
	return &newCustomer, nil
}
func (s *MysqlStore) GetCusotmer(id uint) (*models.Customer, error) {
	customer := &models.Customer{}
	err := s.conn.Model(customer).Where("id=?", id).Error
	if err != nil {
		return nil, err
	}
	return customer, nil

}
func (s *MysqlStore) DeleteCusotmer(id uint) error {
	return s.conn.Model(&models.Customer{}).Delete("id=?", id).Error
}
