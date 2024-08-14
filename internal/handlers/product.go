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
	gin.GET("/products/:id", ph.getProduct)
	gin.POST("/products", ph.createProduct)
	gin.PUT("/products/:id", ph.updateProduct)
	gin.DELETE("/products/:id", ph.deleteProduct)
}
func (ph *productHandler) getProduct(ctx *gin.Context) {
	var productUri dto.GetProductRequest
	if err := ctx.ShouldBindUri(&productUri); err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: err.Error()})
		return
	}

	product, err := ph.store.GetProduct(productUri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Product not found."})
		return
	}

	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: product})

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
		slog.Error("Unable to create product", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to create product"})
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: result})
}
func (ph *productHandler) updateProduct(ctx *gin.Context) {
	var productUri dto.GetProductRequest
	if err := ctx.ShouldBindUri(&productUri); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// check if product exist
	if _, err := ph.store.GetProduct(productUri.ID); err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Product doesn't exist"})
		return
	}
	var productRequest dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// update product
	newProduct := &models.Product{
		Name:     productRequest.Name,
		BarCode:  []byte(productRequest.BarCode),
		Price:    productRequest.Price,
		Qunatity: productRequest.Quantity,
	}
	result, err := ph.store.UpdateProduct(productUri.ID, newProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to update customer"})
		return
	}
	ctx.JSON(http.StatusOK, result)

}
func (ph *productHandler) deleteProduct(ctx *gin.Context) {}
