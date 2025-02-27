package services

import (
	"context"
	"strings"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	accountnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/utilities/account_number.go"
)

type CreateAccount interface {
	CreateAccount(ctx context.Context, req *featuresdtos.CreateAccountRequest) error
}

func NewCreateAccount(
	readAccountsByMobileNumberRepository repositories.ReadAccountsByMobileNumber,
	createAccountRepository repositories.CreateAccount,
	uuidGenerator uuid.UUIDGenerator,
	accountNumberGenerator accountnumberutilities.AccountNumberGenerator,
) CreateAccount {
	return &createAccountImpl{
		readAccountsByMobileNumberRepository: readAccountsByMobileNumberRepository,
		createAccountRepository:              createAccountRepository,
		uuidGenerator:                        uuidGenerator,
		accountNumberGenerator:               accountNumberGenerator,
	}
}

type createAccountImpl struct {
	readAccountsByMobileNumberRepository repositories.ReadAccountsByMobileNumber
	createAccountRepository              repositories.CreateAccount
	uuidGenerator                        uuid.UUIDGenerator
	accountNumberGenerator               accountnumberutilities.AccountNumberGenerator
}

func (handle *createAccountImpl) CreateAccount(ctx context.Context, req *featuresdtos.CreateAccountRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateCreateAccountRequest(req); err != nil {
		return err
	}
	systemTime := time.Now()
	accountEntity := entities.Account{
		ID:              handle.uuidGenerator.GenerateUUID(),
		MobileNumber:    req.MobileNumber,
		AccountNumber:   handle.accountNumberGenerator.GenerateAccountNumber(),
		AccountTypeCode: enumsaccounts.AccountType(strings.ToLower(req.AccountType)).ToAccountTypeCode(),
		BranchCode:      enumsbanks.BanksBranch(strings.ToLower(req.Branch)).ToBanksBranchCode(),
		ActiveSwitch:    true,
		CommonHistoryModelWithUpdate: entities.CommonHistoryModelWithUpdate{
			CommonHistoryModel: entities.CommonHistoryModel{
				CreatedAt: systemTime,
			},
			UpdatedAt: systemTime,
		},
	}
	accounts, err := handle.readAccountsByMobileNumberRepository.ReadAccountsByMobileNumber(ctx, req.MobileNumber)
	if err != nil {
		return err
	}
	if accounts.IncludeAccountTypeCode(req.MobileNumber, accountEntity.AccountTypeCode) {
		return customizeerrors.AccountAlreadyExistsError
	}
	return handle.createAccountRepository.CreateAccount(ctx, accountEntity)
}
