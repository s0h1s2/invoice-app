package repositories

import (
	"context"

	"github.com/s0h1s2/invoice-app/internal/models"
)

type UserRepo interface {
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	ChangePassword(ctx context.Context, userId int, newPassword string)
}
