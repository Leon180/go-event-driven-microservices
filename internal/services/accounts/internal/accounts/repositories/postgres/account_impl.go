package repositoriespostgres

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	postgresdbmodels "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/dbmodels"
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

func (handle *createAccountImpl) CreateAccount(ctx context.Context, account dtos.Account) error {
	systemTime := time.Now()
	dbmodel := postgresdbmodels.Account{
		ID:            account.ID,
		MobileNumber:  account.MobileNumber,
		AccountNumber: account.AccountNumber,
		AccountType:   account.AccountType.ToAccountTypeCode(),
		Branch:        account.Branch.ToBanksBranchCode(),
		ActiveSwitch:  account.ActiveSwitch,
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

type getAccountWithHistoryByMobileNumberImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewGetAccountWithHistoryByMobileNumber(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.GetAccountWithHistoryByMobileNumber {
	return &getAccountWithHistoryByMobileNumberImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *getAccountWithHistoryByMobileNumberImpl) GetAccountWithHistoryByMobileNumber(ctx context.Context, mobileNumber string) (*dtos.AccountWithHistory, error) {
	var dbmodel postgresdbmodels.Account
	if err := handle.db.WithContext(ctx).Where("mobile_number = ?", mobileNumber).Limit(1).Find(&dbmodel).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error getting account with history by mobile number")
		return nil, err
	}
	if dbmodel.ID == "" {
		return nil, nil
	}
	entity := dbmodel.ToDTOsWithHistory()
	return &entity, nil
}

type getAccountWithHistoryImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewGetAccountWithHistory(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.GetAccountWithHistory {
	return &getAccountWithHistoryImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *getAccountWithHistoryImpl) GetAccountWithHistory(ctx context.Context, id string) (*dtos.AccountWithHistory, error) {
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

type updateAccountByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewUpdateAccountByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateAccountByID {
	return &updateAccountByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *updateAccountByIDImpl) UpdateAccountByID(ctx context.Context, update dtos.UpdateAccount) error {
	updateMap := update.ToUpdateMap()
	if update.ID == "" || len(updateMap) == 0 {
		return nil
	}
	if err := handle.db.WithContext(ctx).Model(&postgresdbmodels.Account{}).Where("id = ?", update.ID).Updates(updateMap).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error updating account by id")
		return err
	}
	return nil
}

type deleteAccountByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewDeleteAccountByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.DeleteAccountByID {
	return &deleteAccountByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *deleteAccountByIDImpl) DeleteAccountByID(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Delete(&postgresdbmodels.Account{}).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error deleting account by id")
		return err
	}
	return nil
}
