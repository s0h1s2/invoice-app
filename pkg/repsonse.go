package pkg

type ErrorResponse struct {
	Status int `json:"-"`
	Errors any `json:"errors"`
}
type SuccessResponse struct {
	Data any `json:"data"`
}
