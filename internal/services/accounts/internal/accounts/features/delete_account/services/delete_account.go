package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type DeleteAccount interface {
	DeleteAccount(ctx context.Context, req *featuresdtos.DeleteAccountRequest) error
}

func NewDeleteAccount(
	getAccountWithHistoryRepository repositories.GetAccountWithHistory,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) DeleteAccount {
	return &deleteAccountImpl{
		getAccountWithHistoryRepository: getAccountWithHistoryRepository,
		updateAccountByIDRepository:     updateAccountByIDRepository,
	}
}

type deleteAccountImpl struct {
	getAccountWithHistoryRepository repositories.GetAccountWithHistory
	updateAccountByIDRepository     repositories.UpdateAccountByID
}

func (handle *deleteAccountImpl) DeleteAccount(ctx context.Context, req *featuresdtos.DeleteAccountRequest) error {
	if req == nil {
		return nil
	}
	account, err := handle.getAccountWithHistoryRepository.GetAccountWithHistory(ctx, req.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return customizeerrors.AccountNotFoundError
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
