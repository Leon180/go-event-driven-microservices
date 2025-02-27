package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type DeleteAccount interface {
	DeleteAccount(ctx context.Context, req *featuresdtos.DeleteAccountRequest) error
}

func NewDeleteAccount(
	readAccountRepository repositories.ReadAccount,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) DeleteAccount {
	return &deleteAccountImpl{
		readAccountRepository:       readAccountRepository,
		updateAccountByIDRepository: updateAccountByIDRepository,
	}
}

type deleteAccountImpl struct {
	readAccountRepository       repositories.ReadAccount
	updateAccountByIDRepository repositories.UpdateAccountByID
}

func (handle *deleteAccountImpl) DeleteAccount(ctx context.Context, req *featuresdtos.DeleteAccountRequest) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.IDInvalidError
	}
	account, err := handle.readAccountRepository.ReadAccount(ctx, req.ID)
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
	updateAccount := entities.UpdateAccount{
		ID:           account.ID,
		ActiveSwitch: &activeSwitch,
	}
	return handle.updateAccountByIDRepository.UpdateAccountByID(ctx, updateAccount)
}
