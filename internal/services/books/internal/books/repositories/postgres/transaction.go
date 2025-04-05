package repostgresespostgres

import (
	"context"

	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"gorm.io/gorm"
)

type TransactorImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (t TransactorImpl) BeginTx(ctx context.Context) (customizegorm.Transaction, error) {
	return TransactionImpl{
		Db:            t.db.WithContext(ctx).Begin(),
		ContextLogger: t.contextLogger,
	}, nil
}

type TransactionImpl struct {
	Db            *gorm.DB
	ContextLogger contextloggers.ContextLogger
}

func (t TransactionImpl) Commit() error {
	return t.Db.Commit().Error
}

func (t TransactionImpl) Rollback() error {
	return t.Db.Rollback().Error
}
