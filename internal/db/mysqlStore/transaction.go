package mysqlstore

import (
	"errors"

	"github.com/s0h1s2/invoice-app/internal/operations"
)

type mysqlTransaction struct {
	conn *mysqlStore
}

func NewMysqlStoreTransaction(conn *mysqlStore) *mysqlTransaction {
	return &mysqlTransaction{
		conn: conn,
	}
}
func (my *mysqlTransaction) ExecuteInTransaction(operations operations.Operations) error {
	if len(operations) == 0 {
		return errors.New("no operations provided")
	}
	tx := my.conn.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	for _, op := range operations {
		if err := op(); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
