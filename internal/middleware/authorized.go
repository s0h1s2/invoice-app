package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/s0h1s2/invoice-app/internal/config"
	"github.com/s0h1s2/invoice-app/internal/httperror"
	"github.com/s0h1s2/invoice-app/internal/util"
	"github.com/s0h1s2/invoice-app/pkg"
)

func VerifyAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqToken := ctx.GetHeader("Authorization")
		splittedToken := strings.Split(reqToken, " ")
		if len(splittedToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, pkg.ErrorResponse{Errors: "Bad authorization header format"})
			return
		}
		var userClaims util.UserClaims
		token, err := jwt.ParseWithClaims(splittedToken[1], &userClaims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error druing parsing jwt")
			}
			return []byte(config.Config.Jwt.JwtSecretKey), nil
		})
		if err != nil {
			err := httperror.FromError(err)
			ctx.AbortWithStatusJSON(err.Status, err)
			return
		}
		if token.Valid {
			ctx.Set("user", userClaims.UserID)
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Errors: "Invalid authorization"})
	}
}
