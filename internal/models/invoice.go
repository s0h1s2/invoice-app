package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	InvoiceID  string `gorm:"primary_key"`
	Date       time.Time
	CustomerID int
	Customer   Customer
	Total      float32
}

type InvoiceLine struct {
	gorm.Model
}
