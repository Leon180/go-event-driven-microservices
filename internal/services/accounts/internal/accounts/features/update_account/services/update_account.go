package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type UpdateAccount interface {
	UpdateAccount(ctx context.Context, req *featuredtos.UpdateAccountRequest) error
}

type updateAccountImpl struct {
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber
	updateAccountByIDRepository                   repositories.UpdateAccountByID
}

func NewUpdateAccount(
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber,
	updateAccountByIDRepository repositories.UpdateAccountByID,
) UpdateAccount {
	return &updateAccountImpl{
		getAccountWithHistoryByMobileNumberRepository: getAccountWithHistoryByMobileNumberRepository,
		updateAccountByIDRepository:                   updateAccountByIDRepository,
	}
}

func (handle *updateAccountImpl) UpdateAccount(ctx context.Context, req *featuredtos.UpdateAccountRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateUpdateAccountRequest(*req); err != nil {
		return err
	}
	account, err := handle.getAccountWithHistoryByMobileNumberRepository.GetAccountWithHistoryByMobileNumber(ctx, req.MobileNumber)
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
		AccountType:   req.AccountType,
		BranchAddress: req.BranchAddress,
		ActiveSwitch:  req.ActiveSwitch,
	}
	return handle.updateAccountByIDRepository.UpdateAccountByID(ctx, updateAccount)
}
