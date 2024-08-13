package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/pkg"
)

type supplierHandler struct {
	store repositories.Store
}

func NewSupplierHandler(store repositories.Store) *supplierHandler {
	return &supplierHandler{
		store: store,
	}
}
func (sh *supplierHandler) RegisterSupplierRoutes(route gin.IRouter) {
	route.GET("/suppliers/:id", sh.getSupplier)
	route.POST("/suppliers", sh.createSupplier)
	route.PUT("/suppliers/:id", sh.updateSupplier)
	route.DELETE("/suppliers/:id", sh.deleteSupplier)
}

func (sh *supplierHandler) getSupplier(ctx *gin.Context) {
	var supplierUri dto.GetSupplierRequest
	if err := ctx.ShouldBindUri(&supplierUri); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	result, err := sh.store.GetSupplier(supplierUri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Unable to find supplier."})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (sh *supplierHandler) createSupplier(ctx *gin.Context) {
	var payload dto.CreateSupplierRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	supplier := &models.Supplier{
		Name:  payload.Name,
		Phone: payload.Phone,
	}
	result, err := sh.store.CreateSupplier(supplier)
	if err != nil {
		slog.Error("Unable to create supplier %s", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to create supplier"})
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: result})

}
func (sh *supplierHandler) updateSupplier(ctx *gin.Context) {
	var supplierUri dto.GetSupplierRequest
	if err := ctx.ShouldBindUri(&supplierUri); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// find supplier by id
	supplierId := supplierUri.ID

	_, err := sh.store.GetSupplier(supplierId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Supplier doesn't exist."})
		return
	}
	var payload dto.UpdateSupplierRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	supplier := &models.Supplier{
		Name:  payload.Name,
		Phone: payload.Phone,
	}
	result, err := sh.store.UpdateSupplier(supplierId, supplier)
	if err != nil {
		slog.Error("Unable to update supplier %s", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to update supplier"})
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: result})

}
func (sh *supplierHandler) deleteSupplier(ctx *gin.Context) {
	var supplierUri dto.GetSupplierRequest
	if err := ctx.BindUri(&supplierUri); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	err := sh.store.DeleteSupplier(supplierUri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "Supplier deleted"})
}
