package models

import "time"

type Session struct {
	RefreshToken string
	ExpireAt     time.Time
}
