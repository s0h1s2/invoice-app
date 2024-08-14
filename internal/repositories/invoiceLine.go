package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type InvoiceLineRepository interface {
	CreateInvoiceLine(invoice *models.Invoice)
	GetInvoiceLine(invoiceId uint)
	UpdateInvoiceLine(invoiceId uint, invoice *models.Invoice)
	DeleteInvoiceLine(invoiceId uint)
}
