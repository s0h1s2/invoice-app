package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type TokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}
type UpdateUserRequest struct {
	Password string `json:"password" binding:"required,min=8"`
}
