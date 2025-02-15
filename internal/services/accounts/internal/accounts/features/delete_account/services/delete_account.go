package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	customErrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/errors"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type DeleteAccount interface {
	DeleteAccount(ctx context.Context, req *featuredtos.DeleteAccountRequest) error
}

func NewDeleteAccount(
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) DeleteAccount {
	return &deleteAccountImpl{
		getAccountWithHistoryByMobileNumberRepository: getAccountWithHistoryByMobileNumberRepository,
		updateAccountByIDRepository:                   updateAccountByIDRepository,
	}
}

type deleteAccountImpl struct {
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber
	updateAccountByIDRepository                   repositories.UpdateAccountByID
}

func (handle *deleteAccountImpl) DeleteAccount(ctx context.Context, req *featuredtos.DeleteAccountRequest) error {
	if req == nil {
		return nil
	}
	account, err := handle.getAccountWithHistoryByMobileNumberRepository.GetAccountWithHistoryByMobileNumber(ctx, req.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return customErrors.AccountNotFoundError
	}
	if !account.IsActive() {
		return nil
	}
	activeSwitch := false
	updateAccount := dtos.UpdateAccount{
		ID:           account.ID,
		ActiveSwitch: &activeSwitch,
	}
	return handle.updateAccountByIDRepository.UpdateAccountByID(ctx, updateAccount)
}
