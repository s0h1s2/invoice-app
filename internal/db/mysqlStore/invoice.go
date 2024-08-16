package mysqlstore

import (
	"errors"
	"fmt"
	"time"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type invoiceStore struct {
	conn *mysqlStore
}

func NewInvoiceStore(conn *mysqlStore) *invoiceStore {
	return &invoiceStore{
		conn: conn,
	}
}
func (s *invoiceStore) GetInvoice(invoiceID uint) (*models.Invoice, error) {
	invoice := &models.Invoice{}
	err := s.conn.db.Preload("Lines").Preload("Customer").First(invoice, "id=?", invoiceID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, err
	}
	return invoice, nil
}

func (s *invoiceStore) CreateInvoice(invoice *models.Invoice) (*models.Invoice, error) {
	err := s.conn.db.Create(invoice).Error
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
func (s *invoiceStore) GetLastInvoiceByYear(date time.Time) (*models.Invoice, error) {
	invoice := &models.Invoice{}
	year := date.Year()
	err := s.conn.db.Model(&models.Invoice{}).Where("invoice_id LIKE ?", fmt.Sprintf("%%%d%%", year)).Order("created_at desc").Take(invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, err
	}

	return invoice, nil
}
func (s *invoiceStore) UpdateInvoice(invoiceID uint, invoice *models.Invoice) (*models.Invoice, error) {
	invoiceResult := &models.Invoice{}
	result := s.conn.db.Model(invoiceResult).Clauses(clause.Returning{}).Where("id=?", invoiceID).Updates(invoice)
	if err := result.Error; err != nil {
		return nil, err
	}
	return invoiceResult, nil
}
func (s *invoiceStore) DeleteInvoice(invoiceID uint) error {
	err := s.conn.db.Model(&models.Invoice{}).Delete("id=?", invoiceID).Error
	return err
}
