package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID       uint
	Username     string
	RefreshToken string
	ExpireAt     time.Time
}
