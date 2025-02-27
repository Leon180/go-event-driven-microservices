package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/restore_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type RestoreAccount interface {
	RestoreAccount(ctx context.Context, req *featuresdtos.RestoreAccountRequest) error
}

func NewRestoreAccount(
	readAccountRepository repositories.ReadAccount,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) RestoreAccount {
	return &restoreAccountImpl{
		readAccountRepository:       readAccountRepository,
		updateAccountByIDRepository: updateAccountByIDRepository,
	}
}

type restoreAccountImpl struct {
	readAccountRepository       repositories.ReadAccount
	updateAccountByIDRepository repositories.UpdateAccountByID
}

func (handle *restoreAccountImpl) RestoreAccount(ctx context.Context, req *featuresdtos.RestoreAccountRequest) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	account, err := handle.readAccountRepository.ReadAccount(ctx, req.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return customizeerrors.AccountNotFoundError
	}
	if account.IsActive() {
		return nil
	}
	activeSwitch := true
	updateAccount := entities.UpdateAccount{
		ID:           account.ID,
		ActiveSwitch: &activeSwitch,
	}
	return handle.updateAccountByIDRepository.UpdateAccountByID(ctx, updateAccount)
}
