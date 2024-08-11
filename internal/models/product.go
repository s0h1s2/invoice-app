package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	BarCode    []byte `gorm:"size:13"`
	Qunatity   int
	Price      float32
	SupplierID int
	Supplier   Supplier
	Images     []ProductImage
}
type ProductImage struct {
	gorm.Model
	Name      string
	ProductID int
}
