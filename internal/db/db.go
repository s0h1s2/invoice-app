package db

import (
	"errors"
	"log"
	"log/slog"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlStore struct {
	db *gorm.DB
}

func NewMysqlStore(dsn string) *MysqlStore {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to DB:%s", err.Error())
	}
	return &MysqlStore{
		db: conn,
	}

}
func (s *MysqlStore) Init() {
	s.db.AutoMigrate(&models.User{})
	s.db.AutoMigrate(&models.Customer{})
	s.db.AutoMigrate(&models.Supplier{})
	s.db.AutoMigrate(&models.Product{})
	s.db.AutoMigrate(&models.ProductImage{})

}
func (s *MysqlStore) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, "username=?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}
func (s *MysqlStore) CreateUser(username, password string) (*models.User, error) {
	user, _ := s.FindUserByUsername(username)
	if user != nil {
		return nil, repositories.ErrUsernameAlreadyTake
	}
	newUser := models.User{
		Username: username,
		Password: password,
	}
	err := s.db.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
func (s *MysqlStore) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	err := s.db.Create(customer).Error
	if err != nil {
		return nil, repositories.ErrCustomerCreate
	}
	return customer, nil
}

func (s *MysqlStore) UpdateCustomer(customerId uint, customer *models.Customer) (*models.Customer, error) {
	err := s.db.Where("id=?", customerId).Updates(customer).Error
	if err != nil {
		return nil, repositories.ErrCustomerUpdate
	}
	return customer, nil
}
func (s *MysqlStore) GetCusotmer(id uint) (*models.Customer, error) {
	customer := &models.Customer{}
	err := s.db.First(customer, id).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (s *MysqlStore) DeleteCusotmer(id uint) error {
	return s.db.Model(&models.Customer{}).Delete("id=?", id).Error
}
func (s *MysqlStore) CreateProduct(product *models.Product) (*models.Product, error) {
	err := s.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s *MysqlStore) CreateProductImage(image *models.ProductImage) error {
	return s.db.Model(&models.ProductImage{ProductID: image.ProductID}).Create(image).Error
}
func (s *MysqlStore) UpdateProduct(productId uint, product *models.Product) (*models.Product, error) {
	productResult := &models.Product{}
	result := s.db.Model(productResult).Clauses(clause.Returning{}).Where("id=?", productId).Updates(product)
	if err := result.Error; err != nil {
		return nil, err
	}
	return productResult, nil
}
func (s *MysqlStore) GetProduct(productId uint) (*models.Product, error) {
	product := &models.Product{}
	err := s.db.Preload("Supplier").Preload("Images").Find(product, "id=?", productId).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, repositories.ErrNotFound
	} else if err != nil {
		slog.Error("Error while reading product ", "err", err)
		return nil, err
	}
	return product, nil
}
func (s *MysqlStore) DeleteProduct(productId uint) error {
	return nil
}
func (s *MysqlStore) CreateSupplier(supplier *models.Supplier) (*models.Supplier, error) {
	err := s.db.Create(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
func (s *MysqlStore) GetSupplier(supplierId uint) (*models.Supplier, error) {
	supplier := &models.Supplier{}
	err := s.db.First(supplier, supplierId).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
func (s *MysqlStore) UpdateSupplier(supplierId uint, supplier *models.Supplier) (*models.Supplier, error) {
	err := s.db.Where("id=?", supplierId).Updates(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *MysqlStore) DeleteSupplier(supplierId uint) error {
	err := s.db.Model(&models.Supplier{}).Delete("id=?", supplierId).Error
	return err
}
