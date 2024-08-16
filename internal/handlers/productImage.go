package handlers

import (
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/pkg"
)

type productImageHandler struct {
	image   repositories.ProductImageRepository
	product repositories.ProductRepository
}

func NewProductImageHandler(image repositories.ProductImageRepository, product repositories.ProductRepository) *productImageHandler {
	return &productImageHandler{
		image:   image,
		product: product,
	}
}
func (pm *productImageHandler) RegisterProductImageRoutes(route gin.IRouter) {
	route.POST("/products/:id/image", pm.uploadImage)
}

func (pm *productImageHandler) uploadImage(ctx *gin.Context) {
	var productURI dto.GetProductRequest

	if err := ctx.ShouldBindUri(&productURI); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	productID := productURI.ID
	_, err := pm.product.GetProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: "Product wasn't found."})
		return
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	homeDir, _ := os.Getwd()
	uploadPath := path.Join(homeDir, "assets", "uploads")
	images := form.File["images[]"]
	for _, image := range images {
		fileName := filepath.Base(image.Filename)
		if err := ctx.SaveUploadedFile(image, path.Join(uploadPath, fileName)); err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to upload image."})
			return
		}
		if err := pm.image.CreateProductImage(&models.ProductImage{ProductID: productID, Name: fileName}); err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to create product images"})
			return
		}
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: "Images uploaded successfully"})
}
