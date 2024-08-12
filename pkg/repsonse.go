package pkg

type ErrorResponse struct {
	Errors any `json:"errors"`
}
type SuccessResponse struct {
	Data any `json:"data"`
}
