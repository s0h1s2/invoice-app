package mysqlstore

import "github.com/s0h1s2/invoice-app/internal/models"

type productImageStore struct {
	conn *mysqlStore
}

func NewProductImageStore(conn *mysqlStore) *productImageStore {
	return &productImageStore{
		conn: conn,
	}
}
func (s *productImageStore) CreateProductImage(image *models.ProductImage) error {
	return s.conn.db.Model(&models.ProductImage{ProductID: image.ProductID}).Create(image).Error
}
