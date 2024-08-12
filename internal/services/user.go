package services

import (
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.Store
}

func NewUserService(repo repositories.Store) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (u *UserService) LoginUser(authDto dto.AuthRequest) (string, error) {
	user, err := u.repo.FindUserByUsername(authDto.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authDto.Password))
	if err != nil {
		return "", err
	}

	return "token", nil
}
