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

type productHandler struct {
	store repositories.Store
}

func NewProductHandler(store repositories.Store) *productHandler {
	return &productHandler{
		store: store,
	}
}
func (ph *productHandler) RegisterProductRoutes(gin gin.IRouter) {
	gin.POST("/products", ph.createProduct)
}
func (ph *productHandler) createProduct(ctx *gin.Context) {
	var payload dto.CreateProductRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	_, err := ph.store.GetSupplier(payload.SupplierID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: "Supplier doesn't exist."})
		return
	}
	product := models.Product{
		Name:       payload.Name,
		BarCode:    []byte(payload.BarCode),
		Qunatity:   payload.Quantity,
		Price:      payload.Price,
		SupplierID: payload.SupplierID,
	}

	result, err := ph.store.CreateProduct(&product)
	if err != nil {
		slog.Error("Unable to create product", err)
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to create product"})
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: result})
}
