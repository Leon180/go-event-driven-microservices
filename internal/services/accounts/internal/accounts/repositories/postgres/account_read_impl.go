package repositoriespostgres

import (
	"context"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"gorm.io/gorm"
)

type readAccountsByMobileNumberImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadAccountsByMobileNumber(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadAccountsByMobileNumber {
	return &readAccountsByMobileNumberImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readAccountsByMobileNumberImpl) ReadAccountsByMobileNumber(
	ctx context.Context,
	mobileNumber string,
) (entities.Accounts, error) {
	var accounts entities.Accounts
	if err := handle.db.WithContext(ctx).Where("mobile_number = ?", mobileNumber).Find(&accounts).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("error getting account with history by mobile number")
		return nil, err
	}
	return accounts, nil
}

type readAccountByMobileNumberAndAccountTypeImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadAccountByMobileNumberAndAccountType(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadAccountByMobileNumberAndAccountType {
	return &readAccountByMobileNumberAndAccountTypeImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readAccountByMobileNumberAndAccountTypeImpl) ReadAccountByMobileNumberAndAccountType(
	ctx context.Context,
	mobileNumber string,
	accountTypeCode enums.AccountTypeCode,
) (*entities.Account, error) {
	var account entities.Account
	if err := handle.db.WithContext(ctx).Where("mobile_number = ? AND account_type_code = ?", mobileNumber, accountTypeCode).Limit(1).Find(&account).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("error getting account with history by mobile number and account type")
		return nil, err
	}
	if account.ID == "" {
		return nil, nil
	}
	return &account, nil
}

type readAccountImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadAccount(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadAccount {
	return &readAccountImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readAccountImpl) ReadAccount(ctx context.Context, id string) (*entities.Account, error) {
	var account entities.Account
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&account).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error getting account")
		return nil, err
	}
	if account.ID == "" {
		return nil, nil
	}
	return &account, nil
}
