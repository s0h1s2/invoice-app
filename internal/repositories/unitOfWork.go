package repositories

type Operation func() error

type Operations []Operation

type UnitOfWork interface {
	ExecuteInTransaction(Operations) error
}
