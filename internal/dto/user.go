package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type CreateUserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}
