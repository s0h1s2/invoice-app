package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/httperror"
	"github.com/s0h1s2/invoice-app/internal/middleware"
	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/pkg"
)

type invoiceLineHandler struct {
	invoiceLine repositories.InvoiceLineRepository
}

func NewInvoiceLineHandler(invoiceLine repositories.InvoiceLineRepository) *invoiceLineHandler {
	return &invoiceLineHandler{
		invoiceLine: invoiceLine,
	}
}
func (il *invoiceLineHandler) RegisterInvoiceLineRoutes(routes gin.IRouter) {
	routes.GET("/invoice-lines/:id", middleware.VerifyAuth(), il.getInvoiceLine)
	routes.POST("/invoice-lines", middleware.VerifyAuth(), il.createInvoiceLine)
	routes.PUT("/invoice-lines/:id", middleware.VerifyAuth(), il.updateInvoiceLine)
	routes.DELETE("/invoice-lines/:id", middleware.VerifyAuth(), il.deleteInvoiceLine)
}
func (il *invoiceLineHandler) getInvoiceLine(ctx *gin.Context) {
	var payload dto.GetInvoiceLineRequest
	if err := ctx.ShouldBindUri(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	invoiceLine, err := il.invoiceLine.GetInvoiceLine(payload.ID)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: invoiceLine})
}
func (il *invoiceLineHandler) createInvoiceLine(ctx *gin.Context) {
	var payload dto.CreateIvoiceLineRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
	}
	// TODO:check for product if exist
	// TODO: check for invoice if exist
	newInvoiceLine := &models.InvoiceLine{
		ProductID: payload.ProductID,
		InvoiceID: payload.InvoiceID,
		Price:     payload.Price,
		Quantity:  payload.Quanity,
	}
	il.invoiceLine.CreateInvoiceLine(newInvoiceLine)
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: newInvoiceLine})
}
func (il *invoiceLineHandler) updateInvoiceLine(ctx *gin.Context) {
	var invoiceURI dto.GetInvoiceRequest
	if err := ctx.ShouldBindUri(&invoiceURI); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	var payload dto.UpdateIvoiceLineRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// TODO: check for invoice line id

	updateInvoiceLine := &models.InvoiceLine{
		Price:     payload.Price,
		Quantity:  payload.Quanity,
		ProductID: payload.ProductID,
	}
	_, err := il.invoiceLine.UpdateInvoiceLine(invoiceURI.ID, updateInvoiceLine)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "Invoice line updated"})
}
func (il *invoiceLineHandler) deleteInvoiceLine(ctx *gin.Context) {
	var invoiceURI dto.GetInvoiceRequest
	if err := ctx.ShouldBindUri(&invoiceURI); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// TODO: check if invoice line exist?
	err := il.invoiceLine.DeleteInvoiceLine(invoiceURI.ID)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "Invoice line deleted"})
}
