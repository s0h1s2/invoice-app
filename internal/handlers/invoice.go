package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/operations"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/internal/util"
	"github.com/s0h1s2/invoice-app/pkg"
)

type invoiceHandler struct {
	invoice    repositories.InvoiceRepository
	customer   repositories.CustomerRepository
	unitOfWork util.UnitOfWork
}

func NewInvoiceHandler(invoice repositories.InvoiceRepository, customer repositories.CustomerRepository, unitOfWork util.UnitOfWork) *invoiceHandler {
	return &invoiceHandler{
		invoice:    invoice,
		customer:   customer,
		unitOfWork: unitOfWork,
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
func (ih *invoiceHandler) createInvoice(ctx *gin.Context) {
	var payload dto.CreateInvoiceRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// check if customer exist?
	customer, err := ih.customer.GetCusotmer(payload.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Customer not found"})
		return
	}
	ih.unitOfWork.ExecuteInTransaction(operations.Operations{
		func() error {

		},
	})
}
func (ih *invoiceHandler) updateInvoice(ctx *gin.Context) {}
func (ih *invoiceHandler) deleteInvoice(ctx *gin.Context) {}
