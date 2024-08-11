package services

import (
	"context"

	"github.com/s0h1s2/invoice-app/internal/models"
	"github.com/s0h1s2/invoice-app/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (u *UserService) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := u.repo.FindUserByUsername(context.Background(), "shkar")
	if err != nil {
		return nil, nil
	}
	println(user)
	return nil, nil
}
