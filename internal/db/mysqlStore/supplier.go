package mysqlstore

import (
	"github.com/s0h1s2/invoice-app/internal/models"
)

type supplierStore struct {
	conn *mysqlStore
}

func NewMysqlSupplierStore(conn *mysqlStore) *supplierStore {
	return &supplierStore{
		conn: conn,
	}
}

func (s *supplierStore) CreateSupplier(supplier *models.Supplier) (*models.Supplier, error) {
	err := s.conn.db.Create(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
func (s *supplierStore) GetSupplier(supplierId uint) (*models.Supplier, error) {
	supplier := &models.Supplier{}
	err := s.conn.db.First(supplier, supplierId).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
func (s *supplierStore) UpdateSupplier(supplierId uint, supplier *models.Supplier) (*models.Supplier, error) {
	err := s.conn.db.Where("id=?", supplierId).Updates(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *supplierStore) DeleteSupplier(supplierId uint) error {
	err := s.conn.db.Model(&models.Supplier{}).Delete("id=?", supplierId).Error
	return err
}
