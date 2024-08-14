package repositories

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/models"
)

var (
	ErrUsernameAlreadyTake = errors.New("username already taken")
	ErrCustomerCreate      = errors.New("can't create customer")
	ErrCustomerUpdate      = errors.New("can't update customer")
)

type Store interface {
	/// User
	FindUserByUsername(username string) (*models.User, error)
	CreateUser(username, password string) (*models.User, error)

	/// Customer
	CreateCustomer(customer *models.Customer) (*models.Customer, error)
	UpdateCustomer(customerId uint, customer *models.Customer) (*models.Customer, error)
	GetCusotmer(customerId uint) (*models.Customer, error)
	DeleteCusotmer(customerId uint) error
	/// Product
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(productId uint, product *models.Product) (*models.Product, error)
	GetProduct(productId uint) (*models.Product, error)
	DeleteProduct(productId uint) error
	CreateProductImage(image *models.ProductImage) error
	/// Supplier
	CreateSupplier(supplier *models.Supplier) (*models.Supplier, error)
	GetSupplier(supplierId uint) (*models.Supplier, error)
	UpdateSupplier(supplierId uint, supplier *models.Supplier) (*models.Supplier, error)
	DeleteSupplier(id uint) error

	/// Invoice

}
