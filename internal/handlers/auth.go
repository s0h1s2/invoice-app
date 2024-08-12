package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/services"
	"github.com/s0h1s2/invoice-app/pkg"
)

type userHandler struct {
	userService *services.UserService
}

func NewUserHandler(service *services.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}
func (u *userHandler) RegisterAuthRoutes(route gin.IRouter) {
	route.POST("/users/auth", u.login)
	route.POST("/users", u.createUser)
	route.POST("/users/refresh", u.refreshToken)
	route.PUT("/users/:id", u.changePassword)
}
func (u *userHandler) login(ctx *gin.Context) {
	var auth dto.AuthRequest
	if err := ctx.ShouldBindBodyWithJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := u.userService.LoginUser(auth)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": token})
}
func (u *userHandler) createUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	if err := u.userService.RegisterUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "User created sucessfully"})
}
func (u *userHandler) refreshToken(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Hello": 2})
}
func (u *userHandler) changePassword(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Hello": 3})
}
