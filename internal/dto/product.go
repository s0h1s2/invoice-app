package dto

type GetProductRequest struct {
	ID uint `uri:"id"`
}
type CreateProductRequest struct {
	Name       string  `json:"name" binding:"required"`
	Quantity   int     `json:"quantity" binding:"required"`
	BarCode    string  `json:"barcode" binding:"required,max=13"`
	Price      float32 `json:"price" binding:"required"`
	SupplierID uint    `json:"supplierId" binding:"required"`
}
