package dto

type GetInvoiceLineRequest struct {
	ID uint `uri:"id"`
}
type CreateIvoiceLineRequest struct {
	InvoiceID uint    `json:"invoiceID" binding:"required"`
	ProductID uint    `json:"productID" binding:"required"`
	Quanity   int     `json:"quantity" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
}
type UpdateIvoiceLineRequest struct {
	ProductID uint    `json:"productID"`
	Quanity   int     `json:"quantity"`
	Price     float32 `json:"price"`
}
