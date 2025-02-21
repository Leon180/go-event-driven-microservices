package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type CreateAccount interface {
	CreateAccount(ctx context.Context, req *featuredtos.CreateAccountRequest) error
}

func NewCreateAccount(
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber,
	createAccountRepository repositories.CreateAccount,
	uuidGenerator uuid.UUIDGenerator,
) CreateAccount {
	return &createAccountImpl{
		getAccountWithHistoryByMobileNumberRepository: getAccountWithHistoryByMobileNumberRepository,
		createAccountRepository:                       createAccountRepository,
		uuidGenerator:                                 uuidGenerator,
	}
}

type createAccountImpl struct {
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber
	createAccountRepository                       repositories.CreateAccount
	uuidGenerator                                 uuid.UUIDGenerator
}

func (handle *createAccountImpl) CreateAccount(ctx context.Context, req *featuredtos.CreateAccountRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateCreateAccountRequest(req); err != nil {
		return err
	}
	account, err := handle.getAccountWithHistoryByMobileNumberRepository.GetAccountWithHistoryByMobileNumber(ctx, req.MobileNumber)
	if err != nil {
		return err
	}
	if account != nil {
		return customizeerrors.AccountAlreadyExistsError
	}
	accountDTO := dtos.Account{
		ID:            handle.uuidGenerator.GenerateUUID(),
		MobileNumber:  req.MobileNumber,
		AccountNumber: req.AccountNumber,
		AccountType:   req.AccountType,
		BranchAddress: req.BranchAddress,
		ActiveSwitch:  true,
	}
	return handle.createAccountRepository.CreateAccount(ctx, accountDTO)
}
