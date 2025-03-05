package services

import (
	"context"
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type UpdateAccount interface {
	UpdateAccount(ctx context.Context, req *featuresdtos.UpdateAccountRequest) error
}

type updateAccountImpl struct {
	readAccount                 repositories.ReadAccount
	updateAccountByIDRepository repositories.UpdateAccountByID
}

func NewUpdateAccount(
	readAccount repositories.ReadAccount,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) UpdateAccount {
	return &updateAccountImpl{
		readAccount:                 readAccount,
		updateAccountByIDRepository: updateAccountByIDRepository,
	}
}

func (handle *updateAccountImpl) UpdateAccount(ctx context.Context, req *featuresdtos.UpdateAccountRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateUpdateAccountRequest(*req); err != nil {
		return err
	}
	account, err := handle.readAccount.ReadAccount(ctx, req.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return customizeerrors.AccountNotFoundError
	}
	updateAccount := entities.UpdateAccount{
		ID:            account.ID,
		MobileNumber:  req.MobileNumber,
		AccountNumber: req.AccountNumber,
		BranchCode: func() *enums.BanksBranchCode {
			if req.BranchAddress != nil {
				branchCode := enums.BanksBranch(strings.ToLower(*req.BranchAddress)).ToBanksBranchCode()
				return &branchCode
			}
			return nil
		}(),
	}
	updateAccount.RemoveUnchangedFields(*account)
	if updateAccount.NoUpdates() {
		return customizeerrors.AccountNoUpdatesError
	}
	return handle.updateAccountByIDRepository.UpdateAccountByID(ctx, updateAccount)
}
