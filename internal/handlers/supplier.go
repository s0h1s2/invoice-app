package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
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
}
func (sh *supplierHandler) updateSupplier(ctx *gin.Context) {
}
func (sh *supplierHandler) deleteSupplier(ctx *gin.Context) {
}
