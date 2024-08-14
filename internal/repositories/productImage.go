package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type ProductImageRepository interface {
	CreateProductImage(image *models.ProductImage) error
}
