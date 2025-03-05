package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadAccountRepository := mocksrepositories.NewMockReadAccount(ctrl)
	mockUpdateAccountByIDRepository := mocksrepositories.NewMockUpdateAccountByID(ctrl)

	service := NewUpdateAccount(
		mockReadAccountRepository,
		mockUpdateAccountByIDRepository,
	)

	ctx := context.Background()

	invalidMobileNumber := "123456789000"
	validMobileNumber := "1234567890"
	updateMobileNumber := "1234567891"
	invalidBranchAddress := enums.BanksBranchInvalid.ToString()
	validBranchAddress := enums.BanksBranchTaipeiSongshan.ToString()
	updateBranchAddress := enums.BanksBranchTaipeiZhongshan.ToString()

	// Test cases
	tests := []struct {
		name        string
		setup       func()
		req         *featuresdtos.UpdateAccountRequest
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
			req:         &featuresdtos.UpdateAccountRequest{ID: ""},
			expectError: customizeerrors.InvalidIDError,
		},
		{
			name:        "invalid request - invalid mobile number",
			setup:       func() {},
			req:         &featuresdtos.UpdateAccountRequest{ID: "1234567890", MobileNumber: &invalidMobileNumber},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:        "invalid request - invalid branch address",
			setup:       func() {},
			req:         &featuresdtos.UpdateAccountRequest{ID: "1234567890", BranchAddress: &invalidBranchAddress},
			expectError: customizeerrors.InvalidBranchError,
		},
		{
			name: "account not found",
			setup: func() {
				mockReadAccountRepository.EXPECT().ReadAccount(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.UpdateAccountRequest{ID: "123456789000"},
			expectError: customizeerrors.AccountNotFoundError,
		},
		{
			name: "account no updates",
			setup: func() {
				mockReadAccountRepository.EXPECT().ReadAccount(ctx, "1111111111").Return(&entities.Account{
					ID:              "1111111111",
					MobileNumber:    validMobileNumber,
					AccountNumber:   "1111111111",
					AccountTypeCode: enums.AccountTypeSavings.ToAccountTypeCode(),
					BranchCode:      enums.BanksBranch(validBranchAddress).ToBanksBranchCode(),
					ActiveSwitch:    true,
				}, nil).AnyTimes()
			},
			req: &featuresdtos.UpdateAccountRequest{
				ID:            "1111111111",
				MobileNumber:  &validMobileNumber,
				BranchAddress: &validBranchAddress,
			},
			expectError: customizeerrors.AccountNoUpdatesError,
		},
		{
			name: "account successfully update",
			setup: func() {
				mockReadAccountRepository.EXPECT().ReadAccount(ctx, "1234567890").Return(&entities.Account{
					ID:              "1234567890",
					MobileNumber:    validMobileNumber,
					AccountNumber:   "1234567890",
					AccountTypeCode: enums.AccountTypeSavings.ToAccountTypeCode(),
					BranchCode:      enums.BanksBranch(validBranchAddress).ToBanksBranchCode(),
					ActiveSwitch:    true,
				}, nil).AnyTimes()
				mockUpdateAccountByIDRepository.EXPECT().UpdateAccountByID(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req: &featuresdtos.UpdateAccountRequest{
				ID:            "1234567890",
				MobileNumber:  &updateMobileNumber,
				BranchAddress: &updateBranchAddress,
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.UpdateAccount(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
