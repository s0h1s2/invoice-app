package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/invoice-app/internal/config"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/httperror"
	"github.com/s0h1s2/invoice-app/internal/middleware"
	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/internal/util"
	"github.com/s0h1s2/invoice-app/pkg"
)

type userHandler struct {
	user       repositories.UserRepository
	tokenMaker *util.TokenMaker
}

const (
	accessTokenExpireTime  = time.Hour * 1
	refreshTokenExpireTime = 7 * 24 * time.Hour
)

func NewUserHandler(user repositories.UserRepository, tokenMaker *util.TokenMaker) *userHandler {
	return &userHandler{
		user:       user,
		tokenMaker: tokenMaker,
	}
}
func (u *userHandler) RegisterUserRoutes(route gin.IRouter) {
	route.POST("/users/auth", u.login)
	route.POST("/users", u.createUser)
	route.POST("/users/refresh", u.refreshToken)
	route.PUT("/users/me", middleware.VerifyAuth(), u.updateUser)
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
	accessToken, err := u.tokenMaker.GenerateToken(user.ID, user.Username, config.Config.Jwt.JwtSecretKey, time.Now().Add(accessTokenExpireTime))
	refreshToken, err := u.tokenMaker.GenerateToken(user.ID, user.Username, config.Config.Jwt.JwtSecretKey, time.Now().AddDate(0, 0, int(refreshTokenExpireTime)))
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}

	// TODO: hash session to more safety.
	err = u.user.CreateSession(&models.Session{
		UserID:       user.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		ExpireAt:     time.Now().Add(refreshTokenExpireTime),
	})

	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}})
}
func (u *userHandler) createUser(ctx *gin.Context) {
	var payload dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// TODO: check if username exist before create new user
	hashedPassword := util.HashPassword(payload.Password)
	newUser := &models.User{
		Username: payload.Username,
		Password: hashedPassword,
	}
	if _, err := u.user.CreateUser(newUser); err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "User created sucessfully"})
}
func (u *userHandler) refreshToken(ctx *gin.Context) {
	var payload dto.TokenRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	session, err := u.user.GetSession(payload.RefreshToken)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	if session.ExpireAt.Before(time.Now()) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Status: http.StatusUnauthorized,
			Errors: "Unauthroized token",
		})
		return
	}
	// TODO: block current token in redis or database in case where current access token time isn't expired
	newAccessToken, err := u.tokenMaker.GenerateToken(session.UserID, session.Username, config.Config.Jwt.JwtSecretKey, time.Now().Add(accessTokenExpireTime))
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err.Errors)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}
func (u *userHandler) updateUser(ctx *gin.Context) {
	var payload dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: err.Error()})
		return
	}
	// TODO: this is really bad approach to handle this case.
	// https://github.com/gin-gonic/gin/issues/1123
	// There are a couple ideas i need to implement it.
	// Maybe pass *engine struct in handlers be a good idea whenever authenticated user needed.
	//
	userID, _ := ctx.Get("user")
	hashedPassword := util.HashPassword(payload.Password)
	// update password with new hashed password
	err := u.user.UpdateUserPassword(userID.(uint), hashedPassword)
	if err != nil {
		err := httperror.FromError(err)
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Data: "User updated sucessfully"})
}
