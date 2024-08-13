package dto

type GetCustomerRequest struct {
	ID uint `uri:"id" binding:"required"`
}
type CreateCustomerRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Phone     string  `json:"phone" binding:"required"`
	Balance   float32 `json:"balance" binding:"required"`
}

type UpdateCustomerRequest struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Address   string  `json:"address"`
	Phone     string  `json:"phone"`
	Balance   float32 `json:"balance"`
}
