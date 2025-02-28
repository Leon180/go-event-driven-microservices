package services

// import (
// 	"context"
// 	"testing"

// 	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
// 	mocksuuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/mocks"
// 	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
// 	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/dtos"
// 	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories/mocks"
// 	mockscardnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/utilities/card_number/mocks"
// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestCreateCreditCard(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCreateCreditCardRepository := mocksrepositories.NewMockCreateCreditCard(ctrl)
// 	mockReadCreditCardByMobileNumberAndActiveSwitchRepository := mocksrepositories.NewMockReadCreditCardByMobileNumberAndActiveSwitch(ctrl)
// 	mockUUIDGenerator := mocksuuid.NewMockUUIDGenerator(ctrl)
// 	mockCardNumberGenerator := mockscardnumberutilities.NewMockCardNumberGenerator(ctrl)

// 	service := NewCreateCreditCard(
// 		mockCreateCreditCardRepository,
// 		mockReadCreditCardByMobileNumberAndActiveSwitchRepository,
// 		mockUUIDGenerator,
// 		mockCardNumberGenerator,
// 	)

// 	ctx := context.Background()

// 	// Test cases
// 	validTotalLimit := decimal.NewFromInt(100000).String()
// 	inValidTotalLimit := "invalid"
// 	activeSwitch := true

// 	var tests = []struct {
// 		name        string
// 		setup       func()
// 		req         *featuresdtos.CreateCreditCardRequest
// 		expectError customizeerrors.CustomError
// 	}{
// 		{
// 			name:        "nil request",
// 			setup:       func() {},
// 			req:         nil,
// 			expectError: nil,
// 		},
// 		{
// 			name:        "invalid request - invalid mobile number",
// 			setup:       func() {},
// 			req:         &featuresdtos.CreateCreditCardRequest{MobileNumber: "12345678900", TotalLimit: validTotalLimit},
// 			expectError: customizeerrors.InvalidMobileNumberError,
// 		},
// 		{
// 			name:        "invalid request - invalid total limit",
// 			setup:       func() {},
// 			req:         &featuresdtos.CreateCreditCardRequest{MobileNumber: "1234567890", TotalLimit: inValidTotalLimit},
// 			expectError: customizeerrors.InvalidDecimalError,
// 		},
// 		{
// 			name: "successful credit card creation",
// 			setup: func() {
// 				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
// 				mockCardNumberGenerator.EXPECT().GenerateCardNumber().Return("0000111122223333").AnyTimes()
// 				mockReadCreditCardByMobileNumberAndActiveSwitchRepository.EXPECT().ReadCreditCardByMobileNumberAndActiveSwitch(ctx, "1234567890", &activeSwitch).Return(nil, nil).AnyTimes()
// 				mockCreateCreditCardRepository.EXPECT().CreateCreditCard(ctx, gomock.Any()).Return(nil).AnyTimes()
// 			},
// 			req:         &featuresdtos.CreateCreditCardRequest{MobileNumber: "1234567890", TotalLimit: validTotalLimit},
// 			expectError: nil,
// 		},
// 		{
// 			name: "account already exists",
// 			setup: func() {
// 				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
// 				mockCardNumberGenerator.EXPECT().GenerateCardNumber().Return("0000111122223333").AnyTimes()
// 				mockReadCreditCardByMobileNumberAndActiveSwitchRepository.EXPECT().ReadCreditCardByMobileNumberAndActiveSwitch(ctx, "1111111111", &activeSwitch).Return(entities.CreditCards{
// 					entities.CreditCard{
// 						ID:           "1234567890",
// 						MobileNumber: "1111111111",
// 						CardNumber:   "0000111122223333",
// 						TotalLimit:   decimal.NewFromInt(100000),
// 						AmountUsed:   decimal.NewFromInt(0),
// 						ActiveSwitch: true,
// 					},
// 				}, nil).AnyTimes()
// 			},
// 			req:         &featuresdtos.CreateCreditCardRequest{MobileNumber: "1111111111", TotalLimit: validTotalLimit},
// 			expectError: customizeerrors.CardAlreadyExistsError,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup()
// 			err := service.CreateCreditCard(ctx, test.req)
// 			if test.expectError != nil {
// 				assert.Error(t, err)
// 				assert.Equal(t, test.expectError, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
