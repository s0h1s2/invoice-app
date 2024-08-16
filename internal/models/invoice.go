package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	InvoiceID  string `gorm:"column:invoice_id"`
	ate        time.Time
	CustomerID uint
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
