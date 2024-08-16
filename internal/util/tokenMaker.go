package util

import (
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
func (tm *TokenMaker) GenerateToken(id uint, username, key string, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Username: username,
		UserID:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireAt},
		},
	})
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
