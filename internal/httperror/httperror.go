package httperror

import (
	"log/slog"
	"net/http"

	"github.com/s0h1s2/invoice-app/internal/repositories"
	"github.com/s0h1s2/invoice-app/pkg"
)

func FromError(err error) pkg.ErrorResponse {
	switch err {
	case repositories.ErrNotFound:
		{
			return pkg.ErrorResponse{
				Errors: "Resource not found",
				Status: http.StatusNotFound,
			}
		}
	default:
		{
			slog.Error("Error", "error", err)
			return pkg.ErrorResponse{
				Status: http.StatusInternalServerError,
				Errors: "Internal server error",
			}
		}
	}
}
