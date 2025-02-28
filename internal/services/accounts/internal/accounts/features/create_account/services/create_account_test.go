package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	mocksuuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/mocks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories/mocks"
	mocksaccountnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/utilities/account_number/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadAcountsByMobileNumberRepository := mocksrepositories.NewMockReadAccountsByMobileNumber(ctrl)
	mockCreateAccountRepository := mocksrepositories.NewMockCreateAccount(ctrl)
	mockUUIDGenerator := mocksuuid.NewMockUUIDGenerator(ctrl)
	mockAccountNumberGenerator := mocksaccountnumberutilities.NewMockAccountNumberGenerator(ctrl)

	service := NewCreateAccount(
		mockReadAcountsByMobileNumberRepository,
		mockCreateAccountRepository,
		mockUUIDGenerator,
		mockAccountNumberGenerator,
	)

	ctx := context.Background()

	// Test cases

	var tests = []struct {
		name        string
		setup       func()
		req         *featuresdtos.CreateAccountRequest
		expectError customizeerrors.CustomError
	}{
		{
			name:        "nil request",
			setup:       func() {},
			req:         nil,
			expectError: nil,
		},
		{
			name:        "invalid request - invalid mobile number",
			setup:       func() {},
			req:         &featuresdtos.CreateAccountRequest{MobileNumber: "12345678900", AccountType: enums.AccountTypeSavings.ToString(), Branch: enums.BanksBranchTaipeiSongshan.ToString()},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:        "invalid request - invalid account type",
			setup:       func() {},
			req:         &featuresdtos.CreateAccountRequest{MobileNumber: "1234567890", AccountType: "invalid", Branch: enums.BanksBranchTaipeiSongshan.ToString()},
			expectError: customizeerrors.InvalidAccountTypeError,
		},
		{
			name:        "invalid request - invalid branch",
			setup:       func() {},
			req:         &featuresdtos.CreateAccountRequest{MobileNumber: "1234567890", AccountType: enums.AccountTypeSavings.ToString(), Branch: "invalid"},
			expectError: customizeerrors.InvalidBranchError,
		},
		{
			name: "successful account creation",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockAccountNumberGenerator.EXPECT().GenerateAccountNumber().Return("1234567890").AnyTimes()
				mockReadAcountsByMobileNumberRepository.EXPECT().ReadAccountsByMobileNumber(ctx, "1234567890").Return(nil, nil).AnyTimes()
				mockCreateAccountRepository.EXPECT().CreateAccount(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req:         &featuresdtos.CreateAccountRequest{MobileNumber: "1234567890", AccountType: enums.AccountTypeSavings.ToString(), Branch: enums.BanksBranchTaipeiSongshan.ToString()},
			expectError: nil,
		},
		{
			name: "account already exists",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockAccountNumberGenerator.EXPECT().GenerateAccountNumber().Return("1234567890").AnyTimes()
				mockReadAcountsByMobileNumberRepository.EXPECT().ReadAccountsByMobileNumber(ctx, "1111111111").Return(entities.Accounts{
					entities.Account{
						ID:              "1234567890",
						MobileNumber:    "1111111111",
						AccountNumber:   "1234567890",
						AccountTypeCode: enums.AccountTypeSavings.ToAccountTypeCode(),
						BranchCode:      enums.BanksBranchTaipeiSongshan.ToBanksBranchCode(),
						ActiveSwitch:    true,
					},
				}, nil).AnyTimes()
				mockCreateAccountRepository.EXPECT().CreateAccount(ctx, gomock.Any()).Return(customizeerrors.AccountAlreadyExistsError).AnyTimes()
			},
			req:         &featuresdtos.CreateAccountRequest{MobileNumber: "1111111111", AccountType: enums.AccountTypeSavings.ToString(), Branch: enums.BanksBranchTaipeiSongshan.ToString()},
			expectError: customizeerrors.AccountAlreadyExistsError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.CreateAccount(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
