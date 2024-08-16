package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/httperror"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/internal/util"
	"github.com/s0h1s2/invoice-app/pkg"
)

type userHandler struct {
	user repositories.UserRepository
}
type userClaims struct {
	jwt.Claims
	ID       uint   `json:"userID"`
	Username string `json:"username"`
}

func NewUserHandler(user repositories.UserRepository) *userHandler {
	return &userHandler{
		user: user,
	}
}
func (u *userHandler) RegisterAuthRoutes(route gin.IRouter) {
	route.POST("/users/auth", u.login)
	route.POST("/users", u.createUser)
	route.POST("/users/refresh", u.refreshToken)
	route.PUT("/users/:id", u.updateUser)

}

func (u *userHandler) login(ctx *gin.Context) {
	var auth dto.AuthRequest
	if err := ctx.ShouldBindBodyWithJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	user, err := u.user.FindUserByUsername(auth.Username)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	isPasswordCorrect := util.ComapreHashAndPassword(user.Password, auth.Password)
	if !isPasswordCorrect {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Status: http.StatusUnauthorized, Errors: "Invalid crendentials"})
		return
	}
	userClaims := &userClaims{}
	// generate json web token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenStr, err := token.SigningString()
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"accessToken": tokenStr})
}
func (u *userHandler) createUser(ctx *gin.Context) {
	// var user dto.CreateUserRequest
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
	// 	return
	// }
	// if err := u.user.RegisterUser(user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "User created sucessfully"})
}
func (u *userHandler) refreshToken(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Hello": 2})
}
func (u *userHandler) updateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Hello": 3})
}
