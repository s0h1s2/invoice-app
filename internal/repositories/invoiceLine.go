package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type InvoiceLineRepository interface {
	CreateInvoiceLine(invoice *models.InvoiceLine) (*models.InvoiceLine, error)
	GetInvoiceLine(invoiceID uint) (*models.InvoiceLine, error)
	UpdateInvoiceLine(invoiceID uint, invoice *models.InvoiceLine) (*models.InvoiceLine, error)
	DeleteInvoiceLine(invoiceID uint) error
}
