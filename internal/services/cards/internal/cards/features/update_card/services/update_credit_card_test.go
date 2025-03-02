package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadCreditCardRepository := mocksrepositories.NewMockReadCreditCard(ctrl)
	mockUpdateCreditCardByIDRepository := mocksrepositories.NewMockUpdateCreditCardByID(ctrl)

	service := NewUpdateCreditCard(
		mockReadCreditCardRepository,
		mockUpdateCreditCardByIDRepository,
	)

	ctx := context.Background()

	invalidMobileNumber := "123456789000"
	validMobileNumber := "1234567890"
	updateMobileNumber := "1234567891"
	invalidTotalLimit := "invalid"
	validTotalLimit := decimal.NewFromInt(100000).String()
	invalidAmountUsed := "invalid"
	validAmountUsed := decimal.NewFromInt(100).String()
	amountUsedZero := decimal.NewFromInt(0).String()

	// Test cases
	tests := []struct {
		name        string
		setup       func()
		req         *featuresdtos.UpdateCreditCardRequest
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
			req:         &featuresdtos.UpdateCreditCardRequest{ID: ""},
			expectError: customizeerrors.InvalidIDError,
		},
		{
			name:        "invalid request - invalid mobile number",
			setup:       func() {},
			req:         &featuresdtos.UpdateCreditCardRequest{ID: "1234567890", MobileNumber: &invalidMobileNumber},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:        "invalid request - invalid total limit",
			setup:       func() {},
			req:         &featuresdtos.UpdateCreditCardRequest{ID: "1234567890", TotalLimit: &invalidTotalLimit},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name:        "invalid request - invalid amount used",
			setup:       func() {},
			req:         &featuresdtos.UpdateCreditCardRequest{ID: "1234567890", AmountUsed: &invalidAmountUsed},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name: "credit card not found",
			setup: func() {
				mockReadCreditCardRepository.EXPECT().ReadCreditCard(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.UpdateCreditCardRequest{ID: "123456789000"},
			expectError: customizeerrors.CardNotFoundError,
		},
		{
			name: "credit card no updates",
			setup: func() {
				mockReadCreditCardRepository.EXPECT().ReadCreditCard(ctx, "1111111111").Return(&entities.CreditCard{
					ID:           "1111111111",
					MobileNumber: validMobileNumber,
					CardNumber:   "1111111111",
					TotalLimit: func() decimal.Decimal {
						validTotalLimitDecimal, _ := decimal.NewFromString(validTotalLimit)
						return validTotalLimitDecimal
					}(),
					AmountUsed: func() decimal.Decimal {
						validAmountUsedDecimal, _ := decimal.NewFromString(validAmountUsed)
						return validAmountUsedDecimal
					}(),
					ActiveSwitch: true,
				}, nil).AnyTimes()
			},
			req: &featuresdtos.UpdateCreditCardRequest{
				ID:           "1111111111",
				MobileNumber: &validMobileNumber,
				TotalLimit:   &validTotalLimit,
				AmountUsed:   &amountUsedZero,
			},
			expectError: customizeerrors.CardNoUpdatesError,
		},
		{
			name: "credit card successfully updated",
			setup: func() {
				mockReadCreditCardRepository.EXPECT().ReadCreditCard(ctx, "1234567890").Return(&entities.CreditCard{
					ID:           "1234567890",
					MobileNumber: updateMobileNumber,
					CardNumber:   "1234567890",
					TotalLimit: func() decimal.Decimal {
						validTotalLimitDecimal, _ := decimal.NewFromString(validTotalLimit)
						return validTotalLimitDecimal
					}(),
					AmountUsed:   decimal.Zero,
					ActiveSwitch: true,
				}, nil).AnyTimes()
				mockUpdateCreditCardByIDRepository.EXPECT().
					UpdateCreditCardByID(ctx, gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			req:         &featuresdtos.UpdateCreditCardRequest{ID: "1234567890", AmountUsed: &validAmountUsed},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.UpdateCreditCard(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
