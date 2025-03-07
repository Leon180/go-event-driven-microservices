package postgresgorm

import (
	"context"
)

type Transactor[T Transaction] interface {
	BeginTx(ctx context.Context) (T, error)
}

type Transaction interface {
	Commit() error
	Rollback() error
}
