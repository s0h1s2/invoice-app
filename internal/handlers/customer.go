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

type customerHandler struct {
	customer repositories.CustomerRepository
}

func NewCustomerHandler(customer repositories.CustomerRepository) *customerHandler {
	return &customerHandler{
		customer: customer,
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
	result, err := c.customer.GetCusotmer(payload.ID)
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
	newCusotmer := &models.Customer{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Address:   payload.Address,
		Phone:     payload.Phone,
		Balance:   payload.Balance,
	}
	result, err := c.customer.CreateCustomer(newCusotmer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to create customer"})
		return
	}
	ctx.JSON(http.StatusCreated, pkg.SuccessResponse{Data: result})
}
func (c *customerHandler) updateCustomer(ctx *gin.Context) {
	var customerURI dto.GetCustomerRequest
	if err := ctx.ShouldBindUri(&customerURI); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	var payload dto.UpdateCustomerRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	customerID := customerURI.ID
	_, err := c.customer.GetCusotmer(customerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Customer not found."})
		return
	}
	updatedCustomer := &models.Customer{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Address:   payload.Address,
		Phone:     payload.Phone,
		Balance:   payload.Balance,
	}

	result, err := c.customer.UpdateCustomer(customerID, updatedCustomer)
	if err != nil {
		slog.Error("Unable to update customer due %s", "err", err.Error())
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Errors: "Unable to update customer"})
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: result})
}
func (c *customerHandler) deleteCustomer(ctx *gin.Context) {
	var customerURI dto.GetCustomerRequest
	if err := ctx.ShouldBindUri(&customerURI); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	customerID := customerURI.ID
	_, err := c.customer.GetCusotmer(customerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Errors: "Customer not found."})
		return
	}
	err = c.customer.DeleteCusotmer(customerID)
	if err != nil {
		slog.Error("Unable to delete customer %s", "err", err.Error())
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{})
}
