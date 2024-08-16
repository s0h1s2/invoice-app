package dto

type GetInvoiceLineRequest struct {
	ID uint `uri:"id"`
}
type CreateIvoiceLineRequest struct {
	InvoiceID uint    `json:"invoiceID"`
	ProductID uint    `json:"productID"`
	Quanity   int     `json:"quantity"`
	Price     float32 `json:"price"`
}
type UpdateIvoiceLineRequest struct {
	ProductID uint    `json:"productID"`
	Quanity   int     `json:"quantity"`
	Price     float32 `json:"price"`
}
