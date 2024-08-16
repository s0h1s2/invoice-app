package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/models"
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
	routes.GET("/invoices/:id", ih.getInvoice)
	routes.POST("/invoices", ih.createInvoice)
	routes.PUT("/invoices/:id", ih.updateInvoice)
	routes.DELETE("/invoices/:id", ih.deleteInvoice)
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
	var customerBalanceWithTaxRate float32=customer.Balance
	if payload.Total>=util.GetTaxThreshold(){	
		customerBalanceWithTaxRate=customer.Balance*util.GetTaxRate()
	}
	if customerBalanceWithTaxRate < payload.Total {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: "Customer's balance is not sufficient."})
		return
	}
	date, err := time.Parse(time.DateOnly, payload.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: "Invalid date."})
		return
	}
	lastInvoice, err := ih.invoice.GetLastInvoiceByYear(date)
	if err != nil && !errors.Is(err, repositories.ErrNotFound) {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Internal server error"})
		return
	}

	var newInvoiceID string
	if errors.Is(err, repositories.ErrNotFound) {
		newInvoiceID = fmt.Sprintf("%d-0001", date.Year())
	} else {
		sequence := strings.Split(lastInvoice.InvoiceID, "-")[1]
		nextSequence, err := strconv.ParseInt(sequence, 0, 0)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Internal server error"})
			return
		}
		nextSequence++
		newInvoiceID = fmt.Sprintf("%d-%04d", date.Year(), nextSequence)
	}
	newInvoice := &models.Invoice{
		InvoiceID:  newInvoiceID,
		CustomerID: customer.ID,
		Total:      payload.Total,
	}

	newCustomerBalance:=&models.Customer{
		Balance: customerBalance,
	}
	ih.unitOfWork.ExecuteInTransaction(operations.Operations{
		func() error {
			_, err := ih.invoice.CreateInvoice(newInvoice)
			return err
		},func() error {
			
		}
	})
}
func (ih *invoiceHandler) updateInvoice(ctx *gin.Context) {}
func (ih *invoiceHandler) deleteInvoice(ctx *gin.Context) {}
