package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	entities "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCreditCardsByMobileNumberAndActiveSwitch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadLoanByMobileNumberAndActiveSwitchRepository := mocksrepositories.NewMockReadLoanByMobileNumberAndActiveSwitch(ctrl)

	service := NewGetLoansByMobileNumberAndActiveSwitch(
		mockReadLoanByMobileNumberAndActiveSwitchRepository,
	)

	ctx := context.Background()

	// Test cases
	var tests = []struct {
		name         string
		setup        func()
		req          *featuresdtos.GetLoansRequest
		expectResult entities.Loans
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
			req:          &featuresdtos.GetLoansRequest{MobileNumber: "123456789000"},
			expectResult: nil,
			expectError:  customizeerrors.InvalidMobileNumberError,
		},
		{
			name: "loan not found",
			setup: func() {
				mockReadLoanByMobileNumberAndActiveSwitchRepository.EXPECT().ReadLoanByMobileNumberAndActiveSwitch(ctx, "9999999999", gomock.Any()).Return(nil, nil).AnyTimes()
			},
			req:          &featuresdtos.GetLoansRequest{MobileNumber: "9999999999"},
			expectResult: nil,
			expectError:  customizeerrors.LoanNotFoundError,
		},
		{
			name: "loan found",
			setup: func() {
				mockReadLoanByMobileNumberAndActiveSwitchRepository.EXPECT().ReadLoanByMobileNumberAndActiveSwitch(ctx, "1234567890", nil).Return(entities.Loans{
					entities.Loan{
						ID:           "1234567890",
						MobileNumber: "1234567890",
						LoanNumber:   "1234567890",
						TotalAmount:  decimal.NewFromInt(100000),
						PaidAmount:   decimal.NewFromInt(0),
						ActiveSwitch: true,
					},
				}, nil).AnyTimes()
			},
			req: &featuresdtos.GetLoansRequest{MobileNumber: "1234567890"},
			expectResult: entities.Loans{
				entities.Loan{
					ID:           "1234567890",
					MobileNumber: "1234567890",
					LoanNumber:   "1234567890",
					TotalAmount:  decimal.NewFromInt(100000),
					PaidAmount:   decimal.NewFromInt(0),
					ActiveSwitch: true,
				},
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			result, err := service.GetLoansByMobileNumberAndActiveSwitch(ctx, test.req)
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
					assert.Equal(t, account.LoanNumber, result[i].LoanNumber)
					assert.Equal(t, account.LoanTypeCode, result[i].LoanTypeCode)
					assert.Equal(t, account.TotalAmount, result[i].TotalAmount)
					assert.Equal(t, account.PaidAmount, result[i].PaidAmount)
					assert.Equal(t, account.InterestRate, result[i].InterestRate)
					assert.Equal(t, account.Term, result[i].Term)
					assert.Equal(t, account.ActiveSwitch, result[i].ActiveSwitch)
				}
			}
		})
	}
}
