package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/services"
	"github.com/s0h1s2/invoice-app/pkg"
)

type customerHandler struct {
	customerService *services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *customerHandler {
	return &customerHandler{
		customerService: service,
	}
}
func (c *customerHandler) RegisterCustomerRoutes(route gin.IRouter) {
	route.GET("/customers/:id", c.getCustomer)
	route.POST("/customers", c.createCustomer)
	route.PUT("/customers/:id", c.updateCustomer)
	route.DELETE("/customers/:id", c.deleteCustomer)
}
func (c *customerHandler) getCustomer(ctx *gin.Context) {
	var payload dto.GetCustomerRequest
	if err := ctx.ShouldBindUri(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	result, err := c.customerService.FindCustomerById(payload.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Customer not found"})
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: result})
}
func (c *customerHandler) createCustomer(ctx *gin.Context) {
	var payload dto.CreateCustomerRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	result, err := c.customerService.CreateCustomer(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to create customer"})
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: result})
}
func (c *customerHandler) updateCustomer(ctx *gin.Context) {
	var customerUri dto.GetCustomerRequest
	if err := ctx.ShouldBindUri(&customerUri); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	var payload dto.UpdateCustomerRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	customerID := customerUri.ID
	_, err := c.customerService.FindCustomerById(customerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Customer not found."})
		return
	}
	result, err := c.customerService.UpdateCustomer(customerID, payload)
	if err != nil {
		slog.Error("Unable to update customer due %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to update customer"})
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: result})
}
func (c *customerHandler) deleteCustomer(ctx *gin.Context) {
	var customerUri dto.GetCustomerRequest
	if err := ctx.ShouldBindUri(&customerUri); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	customerID := customerUri.ID
	_, err := c.customerService.FindCustomerById(customerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Customer not found."})
		return
	}
	err = c.customerService.DeleteCustomer(customerID)
	if err != nil {
		slog.Error("Unable to delete customer %s", err.Error())
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{})
}
