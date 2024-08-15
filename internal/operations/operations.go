package operations

type Operation func() error

type Operations []Operation
