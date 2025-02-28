package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/delete_customer/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadCustomerRepository := mocksrepositories.NewMockReadCustomer(ctrl)
	mockUpdateCustomerByIDRepository := mocksrepositories.NewMockUpdateCustomerByID(ctrl)

	service := NewDeleteCustomer(
		mockReadCustomerRepository,
		mockUpdateCustomerByIDRepository,
	)

	ctx := context.Background()

	// Test cases
	var tests = []struct {
		name        string
		setup       func()
		req         *featuresdtos.DeleteCustomerRequest
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
			req:         &featuresdtos.DeleteCustomerRequest{ID: ""},
			expectError: customizeerrors.InvalidIDError,
		},
		{
			name: "customer not found",
			setup: func() {
				mockReadCustomerRepository.EXPECT().ReadCustomer(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.DeleteCustomerRequest{ID: "123456789000"},
			expectError: customizeerrors.CustomerNotFoundError,
		},
		{
			name: "customer already deleted",
			setup: func() {
				mockReadCustomerRepository.EXPECT().ReadCustomer(ctx, "1111111111").Return(&entities.Customer{
					ID:           "1111111111",
					MobileNumber: "1111111111",
					Email:        "test@test.com",
					FirstName:    "test",
					LastName:     "test",
					ActiveSwitch: false,
				}, nil).AnyTimes()
			},
			req:         &featuresdtos.DeleteCustomerRequest{ID: "1111111111"},
			expectError: customizeerrors.CustomerAlreadyDeletedError,
		},
		{
			name: "customer successfully deleted",
			setup: func() {
				mockReadCustomerRepository.EXPECT().ReadCustomer(ctx, "1234567890").Return(&entities.Customer{
					ID:           "1234567890",
					MobileNumber: "1234567890",
					Email:        "test@test.com",
					FirstName:    "test",
					LastName:     "test",
					ActiveSwitch: true,
				}, nil).AnyTimes()
				mockUpdateCustomerByIDRepository.EXPECT().UpdateCustomerByID(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req:         &featuresdtos.DeleteCustomerRequest{ID: "1234567890"},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.DeleteCustomer(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
