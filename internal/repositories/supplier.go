package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type SupplierRepository interface {
	CreateSupplier(supplier *models.Supplier) (*models.Supplier, error)
	GetSupplier(supplierId uint) (*models.Supplier, error)
	UpdateSupplier(supplierId uint, supplier *models.Supplier) (*models.Supplier, error)
	DeleteSupplier(id uint) error
}
