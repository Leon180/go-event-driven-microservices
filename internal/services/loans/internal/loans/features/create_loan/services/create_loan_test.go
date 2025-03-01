package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	mocksuuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/mocks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories/mocks"
	mocksloannumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/utilities/loan_number/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCreditCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateLoanRepository := mocksrepositories.NewMockCreateLoan(ctrl)
	mockReadLoanByMobileNumberAndActiveSwitchRepository := mocksrepositories.NewMockReadLoanByMobileNumberAndActiveSwitch(ctrl)
	mockUUIDGenerator := mocksuuid.NewMockUUIDGenerator(ctrl)
	mockLoanNumberGenerator := mocksloannumberutilities.NewMockLoanNumberGenerator(ctrl)

	service := NewCreateLoan(
		mockCreateLoanRepository,
		mockReadLoanByMobileNumberAndActiveSwitchRepository,
		mockUUIDGenerator,
		mockLoanNumberGenerator,
	)

	ctx := context.Background()

	// Test cases
	validTotalAmount := decimal.NewFromInt(100000).String()
	inValidTotalAmount := "invalid"
	validInterestRate := decimal.NewFromFloat(0.03).String()
	inValidInterestRate := "invalid"
	validTerm := 84
	inValidTerm := 0
	activeSwitch := true

	var tests = []struct {
		name        string
		setup       func()
		req         *featuresdtos.CreateLoanRequest
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
			req:         &featuresdtos.CreateLoanRequest{MobileNumber: "12345678900", LoanType: enums.LoanTypeCar, TotalAmount: validTotalAmount, InterestRate: validInterestRate, Term: validTerm},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:        "invalid request - invalid total amount",
			setup:       func() {},
			req:         &featuresdtos.CreateLoanRequest{MobileNumber: "1234567890", LoanType: enums.LoanTypeCar, TotalAmount: inValidTotalAmount, InterestRate: validInterestRate, Term: validTerm},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name:        "invalid request - invalid interest rate",
			setup:       func() {},
			req:         &featuresdtos.CreateLoanRequest{MobileNumber: "1234567890", LoanType: enums.LoanTypeCar, TotalAmount: validTotalAmount, InterestRate: inValidInterestRate, Term: validTerm},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name:        "invalid request - invalid term",
			setup:       func() {},
			req:         &featuresdtos.CreateLoanRequest{MobileNumber: "1234567890", LoanType: enums.LoanTypeCar, TotalAmount: validTotalAmount, InterestRate: validInterestRate, Term: inValidTerm},
			expectError: customizeerrors.LoanTermInvalidError,
		},
		{
			name: "successful loan creation",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockLoanNumberGenerator.EXPECT().GenerateLoanNumber().Return("0000111122223333").AnyTimes()
				mockReadLoanByMobileNumberAndActiveSwitchRepository.EXPECT().ReadLoanByMobileNumberAndActiveSwitch(ctx, "1234567890", &activeSwitch).Return(nil, nil).AnyTimes()
				mockCreateLoanRepository.EXPECT().CreateLoan(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req:         &featuresdtos.CreateLoanRequest{MobileNumber: "1234567890", LoanType: enums.LoanTypeCar, TotalAmount: validTotalAmount, InterestRate: validInterestRate, Term: validTerm},
			expectError: nil,
		},
		{
			name: "loan already exists",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockLoanNumberGenerator.EXPECT().GenerateLoanNumber().Return("0000111122223333").AnyTimes()
				mockReadLoanByMobileNumberAndActiveSwitchRepository.EXPECT().ReadLoanByMobileNumberAndActiveSwitch(ctx, "1111111111", &activeSwitch).Return(entities.Loans{
					entities.Loan{
						ID:           "1234567890",
						MobileNumber: "1111111111",
						LoanNumber:   "0000111122223333",
						TotalAmount:  decimal.NewFromInt(100000),
						PaidAmount:   decimal.NewFromInt(0),
						ActiveSwitch: true,
					},
				}, nil).AnyTimes()
			},
			req:         &featuresdtos.CreateLoanRequest{MobileNumber: "1111111111", LoanType: enums.LoanTypeCar, TotalAmount: validTotalAmount, InterestRate: validInterestRate, Term: validTerm},
			expectError: customizeerrors.LoanAlreadyExistsError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.CreateLoan(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
