package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	RefreshToken string
	ExpireAt     time.Time
}
