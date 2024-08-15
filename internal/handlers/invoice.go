package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/repositories"
)

type invoiceHandler struct {
	invoice repositories.InvoiceRepository
}

func NewInvoiceRepository(invoice repositories.InvoiceRepository) *invoiceHandler {
	return &invoiceHandler{
		invoice: invoice,
	}
}

func (ih *invoiceHandler) RegisterInvoiceHandler(routes gin.IRouter) {
	routes.GET("/invoice/:id", ih.getInvoice)
	routes.POST("/invoice", ih.createInvoice)
	routes.PUT("/invoice/:id", ih.updateInvoice)
	routes.DELETE("/invoice/:id", ih.deleteInvoice)
}
func (ih *invoiceHandler) getInvoice(ctx *gin.Context) {

}
func (ih *invoiceHandler) createInvoice(ctx *gin.Context) {}
func (ih *invoiceHandler) updateInvoice(ctx *gin.Context) {}
func (ih *invoiceHandler) deleteInvoice(ctx *gin.Context) {}
