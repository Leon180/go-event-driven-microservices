package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	entities "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadLoanRepository := mocksrepositories.NewMockReadLoan(ctrl)
	mockUpdateLoanByIDRepository := mocksrepositories.NewMockUpdateLoanByID(ctrl)

	service := NewUpdateLoan(
		mockReadLoanRepository,
		mockUpdateLoanByIDRepository,
	)

	ctx := context.Background()

	invalidMobileNumber := "123456789000"
	validMobileNumber := "1234567890"
	updateMobileNumber := "1234567891"
	invalidTotalAmount := "invalid"
	validTotalAmount := decimal.NewFromInt(100000).String()
	validPaidAmount := decimal.NewFromInt(0).String()
	invalidPaidAmount := "invalid"
	invalidInterestRate := "invalid"
	validInterestRate := decimal.NewFromFloat(0.03).String()
	invalidTerm := 0
	validTerm := 84

	// Test cases
	tests := []struct {
		name        string
		setup       func()
		req         *featuresdtos.UpdateLoanRequest
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
			req:         &featuresdtos.UpdateLoanRequest{ID: ""},
			expectError: customizeerrors.InvalidIDError,
		},
		{
			name:        "invalid request - invalid mobile number",
			setup:       func() {},
			req:         &featuresdtos.UpdateLoanRequest{ID: "1234567890", MobileNumber: &invalidMobileNumber},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:        "invalid request - invalid total amount",
			setup:       func() {},
			req:         &featuresdtos.UpdateLoanRequest{ID: "1234567890", TotalAmount: &invalidTotalAmount},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name:        "invalid request - invalid paid amount",
			setup:       func() {},
			req:         &featuresdtos.UpdateLoanRequest{ID: "1234567890", PaidAmount: &invalidPaidAmount},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name:        "invalid request - invalid interest rate",
			setup:       func() {},
			req:         &featuresdtos.UpdateLoanRequest{ID: "1234567890", InterestRate: &invalidInterestRate},
			expectError: customizeerrors.InvalidDecimalError,
		},
		{
			name:        "invalid request - invalid term",
			setup:       func() {},
			req:         &featuresdtos.UpdateLoanRequest{ID: "1234567890", Term: &invalidTerm},
			expectError: customizeerrors.LoanTermInvalidError,
		},
		{
			name: "loan not found",
			setup: func() {
				mockReadLoanRepository.EXPECT().ReadLoan(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.UpdateLoanRequest{ID: "123456789000"},
			expectError: customizeerrors.LoanNotFoundError,
		},
		{
			name: "loan no updates",
			setup: func() {
				mockReadLoanRepository.EXPECT().ReadLoan(ctx, "1111111111").Return(&entities.Loan{
					ID:           "1111111111",
					MobileNumber: validMobileNumber,
					LoanNumber:   "1111111111",
					TotalAmount: func() decimal.Decimal {
						validTotalAmountDecimal, _ := decimal.NewFromString(validTotalAmount)
						return validTotalAmountDecimal
					}(),
					PaidAmount: func() decimal.Decimal {
						validPaidAmountDecimal, _ := decimal.NewFromString(validPaidAmount)
						return validPaidAmountDecimal
					}(),
					InterestRate: func() decimal.Decimal {
						validInterestRateDecimal, _ := decimal.NewFromString(validInterestRate)
						return validInterestRateDecimal
					}(),
					Term: validTerm,
				}, nil).AnyTimes()
			},
			req: &featuresdtos.UpdateLoanRequest{
				ID:           "1111111111",
				MobileNumber: &validMobileNumber,
				TotalAmount:  &validTotalAmount,
				InterestRate: &validInterestRate,
			},
			expectError: customizeerrors.LoanNoUpdatesError,
		},
		{
			name: "loan successfully updated",
			setup: func() {
				mockReadLoanRepository.EXPECT().ReadLoan(ctx, "1234567890").Return(&entities.Loan{
					ID:           "1234567890",
					MobileNumber: validMobileNumber,
					LoanNumber:   "1234567890",
					TotalAmount: func() decimal.Decimal {
						validTotalAmountDecimal, _ := decimal.NewFromString(validTotalAmount)
						return validTotalAmountDecimal
					}(),
					PaidAmount:   decimal.Zero,
					ActiveSwitch: true,
				}, nil).AnyTimes()
				mockUpdateLoanByIDRepository.EXPECT().UpdateLoanByID(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req: &featuresdtos.UpdateLoanRequest{
				ID:           "1234567890",
				MobileNumber: &updateMobileNumber,
				TotalAmount:  &validTotalAmount,
				InterestRate: &validInterestRate,
				Term:         &validTerm,
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.UpdateLoan(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
