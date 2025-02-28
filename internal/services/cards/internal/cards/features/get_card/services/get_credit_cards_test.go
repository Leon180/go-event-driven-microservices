package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCreditCardsByMobileNumberAndActiveSwitch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadCreditCardByMobileNumberAndActiveSwitchRepository := mocksrepositories.NewMockReadCreditCardByMobileNumberAndActiveSwitch(ctrl)

	service := NewGetCreditCardsByMobileNumberAndActiveSwitch(
		mockReadCreditCardByMobileNumberAndActiveSwitchRepository,
	)

	ctx := context.Background()

	// Test cases
	var tests = []struct {
		name         string
		setup        func()
		req          *featuresdtos.GetCreditCardsRequest
		expectResult entities.CreditCards
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
			name:         "invalid request - invalid mobile number",
			setup:        func() {},
			req:          &featuresdtos.GetCreditCardsRequest{MobileNumber: "123456789000"},
			expectResult: nil,
			expectError:  customizeerrors.InvalidMobileNumberError,
		},
		{
			name: "credit card not found",
			setup: func() {
				mockReadCreditCardByMobileNumberAndActiveSwitchRepository.EXPECT().ReadCreditCardByMobileNumberAndActiveSwitch(ctx, "9999999999", gomock.Any()).Return(nil, nil).AnyTimes()
			},
			req:          &featuresdtos.GetCreditCardsRequest{MobileNumber: "9999999999"},
			expectResult: nil,
			expectError:  customizeerrors.CardNotFoundError,
		},
		{
			name: "credit card found",
			setup: func() {
				mockReadCreditCardByMobileNumberAndActiveSwitchRepository.EXPECT().ReadCreditCardByMobileNumberAndActiveSwitch(ctx, "1234567890", nil).Return(entities.CreditCards{
					entities.CreditCard{
						ID:           "1234567890",
						MobileNumber: "1234567890",
						CardNumber:   "1234567890",
						TotalLimit:   decimal.NewFromInt(100000),
						AmountUsed:   decimal.NewFromInt(0),
						ActiveSwitch: true,
					},
				}, nil).AnyTimes()
			},
			req: &featuresdtos.GetCreditCardsRequest{MobileNumber: "1234567890"},
			expectResult: entities.CreditCards{
				entities.CreditCard{
					ID:           "1234567890",
					MobileNumber: "1234567890",
					CardNumber:   "1234567890",
					TotalLimit:   decimal.NewFromInt(100000),
					AmountUsed:   decimal.NewFromInt(0),
					ActiveSwitch: true,
				},
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			result, err := service.GetCreditCardsByMobileNumberAndActiveSwitch(ctx, test.req)
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
					assert.Equal(t, account.CardNumber, result[i].CardNumber)
					assert.Equal(t, account.TotalLimit, result[i].TotalLimit)
					assert.Equal(t, account.AmountUsed, result[i].AmountUsed)
					assert.Equal(t, account.ActiveSwitch, result[i].ActiveSwitch)
				}
			}
		})
	}
}
