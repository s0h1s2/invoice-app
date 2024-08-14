package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID         string `gorm:"primary_key"`
	Date       time.Time
	CustomerID int
	Customer   Customer
	Total      float32
	Lines      []InvoiceLine
}

type InvoiceLine struct {
	gorm.Model
	Quantity  int
	Price     float32
	ProductID uint
	InvoiceID uint
}
