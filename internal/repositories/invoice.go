package repositories

import (
	"time"

	"github.com/s0h1s2/invoice-app/internal/models"
)

type InvoiceRepository interface {
	GetInvoice(invoiceID uint) (*models.Invoice, error)
	GetLastInvoiceByYear(date time.Time) (*models.Invoice, error)
	CreateInvoice(invoice *models.Invoice) (*models.Invoice, error)
	UpdateInvoice(invoiceID uint, invoice *models.Invoice) (*models.Invoice, error)
	DeleteInvoice(invoiceID uint) error
}
