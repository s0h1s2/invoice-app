package dto

type GetSupplierRequest struct {
	ID uint `uri:"id"`
}
type CreateSupplierRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
type UpdateSupplierRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
