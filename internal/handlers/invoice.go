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
	"github.com/s0h1s2/invoice-app/internal/httperror"
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
	var invoiceURI dto.GetInvoiceRequest
	if err := ctx.ShouldBindUri(&invoiceURI); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	invoice, err := ih.invoice.GetInvoice(invoiceURI.ID)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: invoice})
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
	var customerBalanceWithTaxRate float32 = customer.Balance
	if payload.Total >= util.GetTaxThreshold() {
		customerBalanceWithTaxRate = customer.Balance * util.GetTaxRate()
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
		ctx.JSON(http.StatusInternalServerError, httperror.FromError(err))
		return
	}

	var newInvoiceID string
	if errors.Is(err, repositories.ErrNotFound) {
		newInvoiceID = fmt.Sprintf("%d-0001", date.Year())
	} else {
		sequence := strings.Split(lastInvoice.InvoiceID, "-")[1]
		nextSequence, err := strconv.ParseInt(strings.TrimLeft(sequence, "0"), 0, 0)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httperror.FromError(err))
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

	newCustomerBalance := &models.Customer{
		Balance: customerBalanceWithTaxRate - payload.Total,
	}
	var invoiceResult *models.Invoice
	err = ih.unitOfWork.ExecuteInTransaction(operations.Operations{
		func() error {
			invoiceResult, err = ih.invoice.CreateInvoice(newInvoice)
			return err
		}, func() error {
			_, err := ih.customer.UpdateCustomer(customer.ID, newCustomerBalance)
			return err
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httperror.FromError(err))
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: invoiceResult})
}
func (ih *invoiceHandler) updateInvoice(ctx *gin.Context) {}
func (ih *invoiceHandler) deleteInvoice(ctx *gin.Context) {}
