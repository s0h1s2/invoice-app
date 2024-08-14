package repositories

import "github.com/s0h1s2/invoice-app/internal/models"

type InvoiceRepository interface {
	CreateInvoice(invoice *models.Invoice)
	GetInvoice(invoiceId uint)
	UpdateInvoice(invoiceId uint, invoice *models.Invoice)
	DeleteInvoice(invoiceId uint)
}
