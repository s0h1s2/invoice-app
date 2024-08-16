package mysqlstore

import (
	"errors"
	"log/slog"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productStore struct {
	conn *mysqlStore
}

func NewMysqlProductStore(conn *mysqlStore) *productStore {
	return &productStore{
		conn: conn,
	}
}
func (s *productStore) GetProduct(productId uint) (*models.Product, error) {
	product := &models.Product{}
	err := s.conn.db.Preload("Supplier").Preload("Images").Find(product, "id=?", productId).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, repositories.ErrNotFound
	} else if err != nil {
		slog.Error("Error while reading product ", "err", err)
		return nil, err
	}
	return product, nil
}

func (s *productStore) CreateProduct(product *models.Product) (*models.Product, error) {
	err := s.conn.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s *productStore) UpdateProduct(productId uint, product *models.Product) (*models.Product, error) {
	productResult := &models.Product{}
	result := s.conn.db.Model(productResult).Clauses(clause.Returning{}).Where("id=?", productId).Updates(product)
	if err := result.Error; err != nil {
		return nil, err
	}
	return productResult, nil
}
func (s *productStore) DeleteProduct(productId uint) error {
	err := s.conn.db.Model(&models.Product{}).Delete("id=?", productId).Error
	return err
}
