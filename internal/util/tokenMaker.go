package util

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenMaker struct{}
type claims struct {
	jwt.RegisteredClaims
	UserID   uint
	Username string
}

func NewTokenMaker() *TokenMaker {
	return &TokenMaker{}
}
func (tm *TokenMaker) GenerateToken(id uint, username, key string, expireAt time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Username: username,
		UserID:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireAt},
		},
	})
	tokenStr, err := token.SignedString(key)
	if err != nil {
		slog.Error("Unable to generate token", "err", err)
	}
	return tokenStr
}
