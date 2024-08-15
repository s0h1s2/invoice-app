package dto

import "time"

type GetInvoiceRequest struct {
	ID uint `uri:"id"`
}
type CreateInvoiceRequest struct {
	Date       time.Time `json:"date" binding:"required"`
	CustomerID uint      `json:"customerId" binding:"required"`
	Total      float32   `json:"total" binding:"required"`
}
type UpdateInvoiceRequest struct {
	Date       time.Time `json:"date"`
	CustomerID uint      `json:"customerId"`
	Total      float32   `json:"total"`
}
