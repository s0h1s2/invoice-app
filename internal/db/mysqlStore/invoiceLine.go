package mysqlstore

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"gorm.io/gorm"
)

type invoiceLineStore struct {
	conn *mysqlStore
}

func NewMysqlInvoiceLineStore(conn *mysqlStore) *invoiceLineStore {
	return &invoiceLineStore{
		conn: conn,
	}
}
func (s *invoiceLineStore) GetInvoiceLine(invoiceID uint) (*models.InvoiceLine, error) {
	result := &models.InvoiceLine{}
	err := s.conn.db.Model(&models.InvoiceLine{}).Take(result, "id=?", invoiceID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, err
	}
	return result, nil
}

func (s *invoiceLineStore) CreateInvoiceLine(invoiceLine *models.InvoiceLine) (*models.InvoiceLine, error) {
	err := s.conn.db.Create(invoiceLine).Error
	if err != nil {
		return nil, err
	}
	return invoiceLine, nil
}
func (s *invoiceLineStore) UpdateInvoiceLine(invoiceID uint, invoice *models.Invoice) (*models.InvoiceLine, error) {

	return nil, nil
}
func (s *invoiceLineStore) DeleteInvoiceLine(invoiceID uint) error {
	return nil
}
