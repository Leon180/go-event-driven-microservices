package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"gorm.io/gorm"
)

type createAccountImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateAccount(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateAccount {
	return &createAccountImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createAccountImpl) CreateAccount(ctx context.Context, account entities.Account) error {
	if err := handle.db.WithContext(ctx).Create(&account).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating account")
		return err
	}
	return nil
}

type createAccountsImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateAccounts(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateAccounts {
	return &createAccountsImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createAccountsImpl) CreateAccounts(ctx context.Context, accounts entities.Accounts) error {
	if len(accounts) == 0 {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(&accounts).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating accounts")
		return err
	}
	return nil
}
