package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	entities "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/delete_loan/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCreditCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadLoanRepository := mocksrepositories.NewMockReadLoan(ctrl)
	mockUpdateLoanByIDRepository := mocksrepositories.NewMockUpdateLoanByID(ctrl)

	service := NewDeleteLoan(
		mockReadLoanRepository,
		mockUpdateLoanByIDRepository,
	)

	ctx := context.Background()

	// Test cases
	var tests = []struct {
		name        string
		setup       func()
		req         *featuresdtos.DeleteLoanRequest
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
			req:         &featuresdtos.DeleteLoanRequest{ID: ""},
			expectError: customizeerrors.InvalidIDError,
		},
		{
			name: "loan not found",
			setup: func() {
				mockReadLoanRepository.EXPECT().ReadLoan(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.DeleteLoanRequest{ID: "123456789000"},
			expectError: customizeerrors.LoanNotFoundError,
		},
		{
			name: "loan already deleted",
			setup: func() {
				mockReadLoanRepository.EXPECT().ReadLoan(ctx, "1111111111").Return(&entities.Loan{
					ID:           "1111111111",
					MobileNumber: "1111111111",
					LoanNumber:   "1111111111",
					TotalAmount:  decimal.NewFromInt(100000),
					ActiveSwitch: false,
				}, nil).AnyTimes()
			},
			req:         &featuresdtos.DeleteLoanRequest{ID: "1111111111"},
			expectError: customizeerrors.LoanAlreadyDeletedError,
		},
		{
			name: "loan successfully deleted",
			setup: func() {
				mockReadLoanRepository.EXPECT().ReadLoan(ctx, "1234567890").Return(&entities.Loan{
					ID:           "1234567890",
					MobileNumber: "1234567890",
					LoanNumber:   "1234567890",
					TotalAmount:  decimal.NewFromInt(100000),
					ActiveSwitch: true,
				}, nil).AnyTimes()
				mockUpdateLoanByIDRepository.EXPECT().UpdateLoanByID(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req:         &featuresdtos.DeleteLoanRequest{ID: "1234567890"},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.DeleteLoan(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
