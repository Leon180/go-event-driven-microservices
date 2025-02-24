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
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber,
	createAccountRepository repositories.CreateAccount,
	uuidGenerator uuid.UUIDGenerator,
	accountNumberGenerator accountnumberutilities.AccountNumberGenerator,
) CreateAccount {
	return &createAccountImpl{
		getAccountWithHistoryByMobileNumberRepository: getAccountWithHistoryByMobileNumberRepository,
		createAccountRepository:                       createAccountRepository,
		uuidGenerator:                                 uuidGenerator,
		accountNumberGenerator:                        accountNumberGenerator,
	}
}

type createAccountImpl struct {
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber
	createAccountRepository                       repositories.CreateAccount
	uuidGenerator                                 uuid.UUIDGenerator
	accountNumberGenerator                        accountnumberutilities.AccountNumberGenerator
}

func (handle *createAccountImpl) CreateAccount(ctx context.Context, req *featuresdtos.CreateAccountRequest) error {
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
		AccountNumber: handle.accountNumberGenerator.GenerateAccountNumber(),
		AccountType:   enumsaccounts.AccountType(strings.ToLower(req.AccountType)),
		Branch:        enumsbanks.BanksBranch(strings.ToLower(req.Branch)),
		ActiveSwitch:  true,
	}
	return handle.createAccountRepository.CreateAccount(ctx, accountDTO)
}
