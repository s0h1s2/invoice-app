package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string `gorm:"size:100"`
	LastName  string `gorm:"size:100"`
	Address   string
	Phone     string `gorm:"size:20"`
	Balance   float32
}
