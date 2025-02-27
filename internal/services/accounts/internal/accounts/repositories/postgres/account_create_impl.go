package repositoriespostgres

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	postgresdbmodels "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/dbmodels"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"github.com/samber/lo"
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

func (handle *createAccountImpl) CreateAccount(ctx context.Context, account dtos.Account) error {
	systemTime := time.Now()
	dbmodel := postgresdbmodels.Account{
		ID:              account.ID,
		MobileNumber:    account.MobileNumber,
		AccountNumber:   account.AccountNumber,
		AccountTypeCode: account.AccountTypeCode,
		BranchCode:      account.BranchCode,
		ActiveSwitch:    account.ActiveSwitch,
		CommonHistoryModelWithUpdate: postgresdbmodels.CommonHistoryModelWithUpdate{
			CommonHistoryModel: postgresdbmodels.CommonHistoryModel{
				CreatedAt: systemTime,
			},
			UpdatedAt: systemTime,
		},
	}
	if err := handle.db.WithContext(ctx).Create(&dbmodel).Error; err != nil {
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

func (handle *createAccountsImpl) CreateAccounts(ctx context.Context, accounts []dtos.Account) error {
	systemTime := time.Now()
	dbmodels := lo.Map(accounts, func(account dtos.Account, _ int) postgresdbmodels.Account {
		return postgresdbmodels.Account{
			ID:              account.ID,
			MobileNumber:    account.MobileNumber,
			AccountNumber:   account.AccountNumber,
			AccountTypeCode: account.AccountTypeCode,
			BranchCode:      account.BranchCode,
			ActiveSwitch:    account.ActiveSwitch,
			CommonHistoryModelWithUpdate: postgresdbmodels.CommonHistoryModelWithUpdate{
				CommonHistoryModel: postgresdbmodels.CommonHistoryModel{
					CreatedAt: systemTime,
				},
				UpdatedAt: systemTime,
			},
		}
	})
	if err := handle.db.WithContext(ctx).Create(&dbmodels).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating accounts")
		return err
	}
	return nil
}
