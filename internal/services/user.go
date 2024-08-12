package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/s0h1s2/invoice-app/internal/config"
	"github.com/s0h1s2/invoice-app/internal/dto"
	"github.com/s0h1s2/invoice-app/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userClaim struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}
type UserService struct {
	repo repositories.Store
}

func NewUserService(repo repositories.Store) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (u *UserService) RegisterUser(createUserDto dto.CreateUserDto) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, e := u.repo.CreateUser(createUserDto.Username, string(hashedPassword))
	return e
}
func (u *UserService) LoginUser(authDto dto.AuthRequest) (*string, error) {
	user, err := u.repo.FindUserByUsername(authDto.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authDto.Password))
	if err != nil {
		return nil, err
	}
	userClaim := userClaim{
		Username: user.Username,
		ID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "issuer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, userClaim)
	tokenStr, err := token.SignedString(config.Config.Jwt.JwtSecretKey)
	if err != nil {
		return nil, err
	}
	return &tokenStr, nil
}
