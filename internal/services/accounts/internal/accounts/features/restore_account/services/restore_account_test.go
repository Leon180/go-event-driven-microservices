package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/restore_account/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRestoreAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadAccountRepository := mocksrepositories.NewMockReadAccount(ctrl)
	mockUpdateAccountByIDRepository := mocksrepositories.NewMockUpdateAccountByID(ctrl)

	service := NewRestoreAccount(
		mockReadAccountRepository,
		mockUpdateAccountByIDRepository,
	)

	ctx := context.Background()

	// Test cases
	var tests = []struct {
		name        string
		setup       func()
		req         *featuresdtos.RestoreAccountRequest
		expectError customizeerrors.CustomError
	}{
		{
			name:        "nil request",
			setup:       func() {},
			req:         nil,
			expectError: nil,
		},
		{
			name:        "invalid request - empty id",
			setup:       func() {},
			req:         &featuresdtos.RestoreAccountRequest{ID: ""},
			expectError: customizeerrors.IDInvalidError,
		},
		{
			name: "account not found",
			setup: func() {
				mockReadAccountRepository.EXPECT().ReadAccount(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.RestoreAccountRequest{ID: "123456789000"},
			expectError: customizeerrors.AccountNotFoundError,
		},
		{
			name: "account already restored",
			setup: func() {
				mockReadAccountRepository.EXPECT().ReadAccount(ctx, "1111111111").Return(&entities.Account{
					ID:              "1111111111",
					MobileNumber:    "1111111111",
					AccountNumber:   "1111111111",
					AccountTypeCode: enumsaccounts.AccountTypeSavings.ToAccountTypeCode(),
					BranchCode:      enumsbanks.BanksBranchTaipeiSongshan.ToBanksBranchCode(),
					ActiveSwitch:    true,
				}, nil).AnyTimes()
			},
			req:         &featuresdtos.RestoreAccountRequest{ID: "1111111111"},
			expectError: nil,
		},
		{
			name: "account successfully restored",
			setup: func() {
				mockReadAccountRepository.EXPECT().ReadAccount(ctx, "1234567890").Return(&entities.Account{
					ID:              "1234567890",
					MobileNumber:    "1234567890",
					AccountNumber:   "1234567890",
					AccountTypeCode: enumsaccounts.AccountTypeSavings.ToAccountTypeCode(),
					BranchCode:      enumsbanks.BanksBranchTaipeiSongshan.ToBanksBranchCode(),
					ActiveSwitch:    false,
				}, nil).AnyTimes()
				mockUpdateAccountByIDRepository.EXPECT().UpdateAccountByID(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req:         &featuresdtos.RestoreAccountRequest{ID: "1234567890"},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.RestoreAccount(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
