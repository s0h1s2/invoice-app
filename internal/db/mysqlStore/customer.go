package mysqlstore

import (
	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
)

type customerStore struct {
	conn *mysqlStore
}

func NewMysqlCustomerStore(conn *mysqlStore) *customerStore {
	return &customerStore{
		conn: conn,
	}
}
func (s *customerStore) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	err := s.conn.db.Create(customer).Error
	if err != nil {
		return nil, repositories.ErrCustomerCreate
	}
	return customer, nil
}

func (s *customerStore) UpdateCustomer(customerId uint, customer *models.Customer) (*models.Customer, error) {
	err := s.conn.db.Where("id=?", customerId).Updates(customer).Error
	if err != nil {
		return nil, repositories.ErrCustomerUpdate
	}
	return customer, nil
}
func (s *customerStore) GetCusotmer(id uint) (*models.Customer, error) {
	customer := &models.Customer{}
	err := s.conn.db.First(customer, id).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (s *customerStore) DeleteCusotmer(id uint) error {
	return s.conn.db.Model(&models.Customer{}).Delete("id=?", id).Error
}
