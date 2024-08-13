package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name     string
	Phone    string `gorm:"size:20"`
	Products []Product
}
