package postgresgorm

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/dbmodels"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"gorm.io/gorm"
)

type createAccountImpl struct {
	db *gorm.DB
}

func NewCreateAccount(db *gorm.DB) repositories.CreateAccount {
	return &createAccountImpl{db: db}
}

func (handle *createAccountImpl) CreateAccount(ctx context.Context, account dtos.Account) error {
	systemTime := time.Now()
	dbmodel := dbmodels.Account{
		ID:            account.ID,
		MobileNumber:  account.MobileNumber,
		AccountNumber: account.AccountNumber,
		AccountType:   account.AccountType,
		BranchAddress: account.BranchAddress,
		ActiveSwitch:  account.ActiveSwitch,
		CommonHistoryModelWithUpdate: dbmodels.CommonHistoryModelWithUpdate{
			CommonHistoryModel: dbmodels.CommonHistoryModel{
				CreatedAt: systemTime,
			},
			UpdatedAt: systemTime,
		},
	}
	if err := handle.db.WithContext(ctx).Create(&dbmodel).Error; err != nil {
		utilities.LogError(ctx, "error creating account: %v", err)
		return err
	}
	return nil
}

type getAccountWithHistoryByMobileNumberImpl struct {
	db *gorm.DB
}

func NewGetAccountWithHistoryByMobileNumber(db *gorm.DB) repositories.GetAccountWithHistoryByMobileNumber {
	return &getAccountWithHistoryByMobileNumberImpl{db: db}
}

func (handle *getAccountWithHistoryByMobileNumberImpl) GetAccountWithHistoryByMobileNumber(ctx context.Context, mobileNumber string) (*dtos.AccountWithHistory, error) {
	var dbmodel dbmodels.Account
	if err := handle.db.WithContext(ctx).Where("mobile_number = ?", mobileNumber).Limit(1).Find(&dbmodel).Error; err != nil {
		utilities.LogError(ctx, "error getting account with history by mobile number: %v", err)
		return nil, err
	}
	if dbmodel.ID == "" {
		return nil, nil
	}
	entity := dbmodel.ToDTOsWithHistory()
	return &entity, nil
}

type updateAccountByIDImpl struct {
	db *gorm.DB
}

func NewUpdateAccountByID(db *gorm.DB) repositories.UpdateAccountByID {
	return &updateAccountByIDImpl{db: db}
}

func (handle *updateAccountByIDImpl) UpdateAccountByID(ctx context.Context, update dtos.UpdateAccount) error {
	updateMap := update.ToUpdateMap()
	if update.ID == "" || len(updateMap) == 0 {
		return nil
	}
	if err := handle.db.WithContext(ctx).Model(&dbmodels.Account{}).Where("id = ?", update.ID).Updates(updateMap).Error; err != nil {
		utilities.LogError(ctx, "error updating account by id: %v", err)
		return err
	}
	return nil
}

type deleteAccountByIDImpl struct {
	db *gorm.DB
}

func NewDeleteAccountByID(db *gorm.DB) repositories.DeleteAccountByID {
	return &deleteAccountByIDImpl{db: db}
}

func (handle *deleteAccountByIDImpl) DeleteAccountByID(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Delete(&dbmodels.Account{}).Error; err != nil {
		utilities.LogError(ctx, "error deleting account by id: %v", err)
		return err
	}
	return nil
}
