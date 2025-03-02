package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccountsByMobileNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadAccountsByMobileNumberRepository := mocksrepositories.NewMockReadAccountsByMobileNumber(ctrl)

	service := NewGetAccountsByMobileNumber(
		mockReadAccountsByMobileNumberRepository,
	)

	ctx := context.Background()

	// Test cases
	tests := []struct {
		name         string
		setup        func()
		req          *featuresdtos.GetAccountsByMobileNumberRequest
		expectResult entities.Accounts
		expectError  customizeerrors.CustomError
	}{
		{
			name:         "nil request",
			setup:        func() {},
			req:          nil,
			expectResult: nil,
			expectError:  nil,
		},
		{
			name:         "invalid request - empty mobile number",
			setup:        func() {},
			req:          &featuresdtos.GetAccountsByMobileNumberRequest{MobileNumber: ""},
			expectResult: nil,
			expectError:  customizeerrors.InvalidMobileNumberError,
		},
		{
			name:         "invalid request - invalid mobile number",
			setup:        func() {},
			req:          &featuresdtos.GetAccountsByMobileNumberRequest{MobileNumber: "123456789000"},
			expectResult: nil,
			expectError:  customizeerrors.InvalidMobileNumberError,
		},
		{
			name: "account not found",
			setup: func() {
				mockReadAccountsByMobileNumberRepository.EXPECT().
					ReadAccountsByMobileNumber(ctx, "9999999999").
					Return(nil, nil).
					AnyTimes()
			},
			req:          &featuresdtos.GetAccountsByMobileNumberRequest{MobileNumber: "9999999999"},
			expectResult: nil,
			expectError:  customizeerrors.AccountNotFoundError,
		},
		{
			name: "account found",
			setup: func() {
				mockReadAccountsByMobileNumberRepository.EXPECT().
					ReadAccountsByMobileNumber(ctx, "1234567890").
					Return(entities.Accounts{
						entities.Account{
							ID:              "1234567890",
							MobileNumber:    "1234567890",
							AccountNumber:   "1234567890",
							AccountTypeCode: enums.AccountTypeSavings.ToAccountTypeCode(),
							BranchCode:      enums.BanksBranchTaipeiSongshan.ToBanksBranchCode(),
							ActiveSwitch:    true,
						},
					}, nil).
					AnyTimes()
			},
			req: &featuresdtos.GetAccountsByMobileNumberRequest{MobileNumber: "1234567890"},
			expectResult: entities.Accounts{
				entities.Account{
					ID:              "1234567890",
					MobileNumber:    "1234567890",
					AccountNumber:   "1234567890",
					AccountTypeCode: enums.AccountTypeSavings.ToAccountTypeCode(),
					BranchCode:      enums.BanksBranchTaipeiSongshan.ToBanksBranchCode(),
					ActiveSwitch:    true,
				},
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			result, err := service.GetAccountsByMobileNumber(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectResult, result)
			}
			if test.expectResult != nil {
				assert.Equal(t, len(test.expectResult), len(result))
				for i, account := range test.expectResult {
					assert.Equal(t, account.ID, result[i].ID)
					assert.Equal(t, account.MobileNumber, result[i].MobileNumber)
					assert.Equal(t, account.AccountNumber, result[i].AccountNumber)
					assert.Equal(t, account.AccountTypeCode, result[i].AccountTypeCode)
					assert.Equal(t, account.BranchCode, result[i].BranchCode)
					assert.Equal(t, account.ActiveSwitch, result[i].ActiveSwitch)
				}
			}
		})
	}
}
