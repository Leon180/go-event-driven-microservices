package services

import (
	"context"
	"strings"

	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	accountnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/utilities/account_number.go"
)

type CreateAccount interface {
	CreateAccount(ctx context.Context, req *featuresdtos.CreateAccountRequest) error
}

func NewCreateAccount(
	readAccountsWithHistoryByMobileNumberRepository repositories.ReadAccountsWithHistoryByMobileNumber,
	createAccountRepository repositories.CreateAccount,
	uuidGenerator uuid.UUIDGenerator,
	accountNumberGenerator accountnumberutilities.AccountNumberGenerator,
) CreateAccount {
	return &createAccountImpl{
		readAccountsWithHistoryByMobileNumberRepository: readAccountsWithHistoryByMobileNumberRepository,
		createAccountRepository:                         createAccountRepository,
		uuidGenerator:                                   uuidGenerator,
		accountNumberGenerator:                          accountNumberGenerator,
	}
}

type createAccountImpl struct {
	readAccountsWithHistoryByMobileNumberRepository repositories.ReadAccountsWithHistoryByMobileNumber
	createAccountRepository                         repositories.CreateAccount
	uuidGenerator                                   uuid.UUIDGenerator
	accountNumberGenerator                          accountnumberutilities.AccountNumberGenerator
}

func (handle *createAccountImpl) CreateAccount(ctx context.Context, req *featuresdtos.CreateAccountRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateCreateAccountRequest(req); err != nil {
		return err
	}
	accountDTO := dtos.Account{
		ID:              handle.uuidGenerator.GenerateUUID(),
		MobileNumber:    req.MobileNumber,
		AccountNumber:   handle.accountNumberGenerator.GenerateAccountNumber(),
		AccountTypeCode: enumsaccounts.AccountType(strings.ToLower(req.AccountType)).ToAccountTypeCode(),
		BranchCode:      enumsbanks.BanksBranch(strings.ToLower(req.Branch)).ToBanksBranchCode(),
		ActiveSwitch:    true,
	}
	accounts, err := handle.readAccountsWithHistoryByMobileNumberRepository.ReadAccountsWithHistoryByMobileNumber(ctx, req.MobileNumber)
	if err != nil {
		return err
	}
	if dtos.AccountsWithHistory(accounts).IncludeAccountTypeCode(req.MobileNumber, accountDTO.AccountTypeCode) {
		return customizeerrors.AccountAlreadyExistsError
	}
	return handle.createAccountRepository.CreateAccount(ctx, accountDTO)
}
