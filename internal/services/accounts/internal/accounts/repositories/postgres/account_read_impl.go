package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	postgresdbmodels "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/dbmodels"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type readAccountsWithHistoryByMobileNumberImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadAccountsWithHistoryByMobileNumber(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadAccountsWithHistoryByMobileNumber {
	return &readAccountsWithHistoryByMobileNumberImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readAccountsWithHistoryByMobileNumberImpl) ReadAccountsWithHistoryByMobileNumber(ctx context.Context, mobileNumber string) ([]dtos.AccountWithHistory, error) {
	var dbmodels []postgresdbmodels.Account
	if err := handle.db.WithContext(ctx).Where("mobile_number = ?", mobileNumber).Find(&dbmodels).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error getting account with history by mobile number")
		return nil, err
	}
	entities := lo.Map(dbmodels, func(dbmodel postgresdbmodels.Account, _ int) dtos.AccountWithHistory {
		return dbmodel.ToDTOsWithHistory()
	})
	return entities, nil
}

type readAccountWithHistoryByMobileNumberAndAccountTypeImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadAccountWithHistoryByMobileNumberAndAccountType(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadAccountWithHistoryByMobileNumberAndAccountType {
	return &readAccountWithHistoryByMobileNumberAndAccountTypeImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readAccountWithHistoryByMobileNumberAndAccountTypeImpl) ReadAccountWithHistoryByMobileNumberAndAccountType(ctx context.Context, mobileNumber string, accountTypeCode enumsaccounts.AccountTypeCode) (*dtos.AccountWithHistory, error) {
	var dbmodels postgresdbmodels.Account
	if err := handle.db.WithContext(ctx).Where("mobile_number = ? AND account_type_code = ?", mobileNumber, accountTypeCode).Limit(1).Find(&dbmodels).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error getting account with history by mobile number and account type")
		return nil, err
	}
	if dbmodels.ID == "" {
		return nil, nil
	}
	entity := dbmodels.ToDTOsWithHistory()
	return &entity, nil
}

type readAccountWithHistoryImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadAccountWithHistory(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadAccountWithHistory {
	return &readAccountWithHistoryImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readAccountWithHistoryImpl) ReadAccountWithHistory(ctx context.Context, id string) (*dtos.AccountWithHistory, error) {
	var dbmodel postgresdbmodels.Account
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&dbmodel).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error getting account with history")
		return nil, err
	}
	if dbmodel.ID == "" {
		return nil, nil
	}
	entity := dbmodel.ToDTOsWithHistory()
	return &entity, nil
}
