package util

import (
	"github.com/s0h1s2/invoice-app/internal/operations"
)

type UnitOfWork interface {
	ExecuteInTransaction(operations.Operations) error
}
