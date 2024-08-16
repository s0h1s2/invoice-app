package mysqlstore

import (
	"log"

	"github.com/s0h1s2/invoice-app/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlStore struct {
	db *gorm.DB
}

func NewMysqlStore(dsn string) *mysqlStore {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to DB:%s", err.Error())
	}
	return &mysqlStore{
		db: conn,
	}

}
func (s *mysqlStore) Init() {
	s.db.AutoMigrate(&models.User{})
	s.db.AutoMigrate(&models.Session{})
	s.db.AutoMigrate(&models.Customer{})
	s.db.AutoMigrate(&models.Supplier{})
	s.db.AutoMigrate(&models.Product{})
	s.db.AutoMigrate(&models.ProductImage{})
	s.db.AutoMigrate(&models.Invoice{})
	s.db.AutoMigrate(&models.InvoiceLine{})
}
func (s *mysqlStore) GetDB() *gorm.DB {
	return s.db
}
