package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(productId uint, product *models.Product) (*models.Product, error)
	GetProduct(productId uint) (*models.Product, error)
	DeleteProduct(productId uint) error
}
