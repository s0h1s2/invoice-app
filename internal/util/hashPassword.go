package util

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashCost = bcrypt.DefaultCost
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		slog.Error("Error during hash password", "err", err)
		return ""
	}
	return string(hashedPassword)
}

func ComapreHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}
	return true
}
