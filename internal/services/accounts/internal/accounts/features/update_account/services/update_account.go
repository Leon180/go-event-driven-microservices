package services

import (
	"context"
	"strings"

	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type UpdateAccount interface {
	UpdateAccount(ctx context.Context, id string, req *featuresdtos.UpdateAccountRequest) error
}

type updateAccountImpl struct {
	readAccountWithHistory      repositories.ReadAccountWithHistory
	updateAccountByIDRepository repositories.UpdateAccountByID
}

func NewUpdateAccount(
	readAccountWithHistory repositories.ReadAccountWithHistory,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) UpdateAccount {
	return &updateAccountImpl{
		readAccountWithHistory:      readAccountWithHistory,
		updateAccountByIDRepository: updateAccountByIDRepository,
	}
}

func (handle *updateAccountImpl) UpdateAccount(ctx context.Context, id string, req *featuresdtos.UpdateAccountRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateUpdateAccountRequest(*req); err != nil {
		return err
	}
	account, err := handle.readAccountWithHistory.ReadAccountWithHistory(ctx, id)
	if err != nil {
		return err
	}
	if account == nil {
		return customizeerrors.AccountNotFoundError
	}
	updateAccount := dtos.UpdateAccount{
		ID:            account.ID,
		MobileNumber:  req.MobileNumber,
		AccountNumber: req.AccountNumber,
		BranchCode: func() *enumsbanks.BanksBranchCode {
			if req.BranchAddress != nil {
				branchCode := enumsbanks.BanksBranch(strings.ToLower(*req.BranchAddress)).ToBanksBranchCode()
				return &branchCode
			}
			return nil
		}(),
	}
	updateAccount.RemoveUnchangedFields(account.Account)
	if updateAccount.NoUpdates() {
		return customizeerrors.AccountNoUpdatesError
	}
	return handle.updateAccountByIDRepository.UpdateAccountByID(ctx, updateAccount)
}
