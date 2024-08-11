package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/services"
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
	route.POST("/users/refresh", u.refreshToken)
	route.PUT("/users/:id", u.changePassword)
}
func (u *userHandler) login(ctx *gin.Context) {
	var auth dto.AuthRequest
	if err := ctx.ShouldBindBodyWithJSON(&auth); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	// user, err := u.userService.FindUserByUsername(context.Background(), auth.Username)

	ctx.JSON(200, gin.H{"Hello": 1})
}
func (u *userHandler) refreshToken(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Hello": 2})
}
func (u *userHandler) changePassword(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Hello": 3})
}
