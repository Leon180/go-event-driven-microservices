package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	mocksuuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/mocks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateCustomerRepository := mocksrepositories.NewMockCreateCustomer(ctrl)
	mockReadCustomerByMobileNumberAndActiveSwitchRepository := mocksrepositories.NewMockReadCustomerByMobileNumberAndActiveSwitch(
		ctrl,
	)
	mockUUIDGenerator := mocksuuid.NewMockUUIDGenerator(ctrl)

	service := NewCreateCustomer(
		mockCreateCustomerRepository,
		mockReadCustomerByMobileNumberAndActiveSwitchRepository,
		mockUUIDGenerator,
	)

	ctx := context.Background()

	// Test cases

	tests := []struct {
		name        string
		setup       func()
		req         *featuresdtos.CreateCustomerRequest
		expectError customizeerrors.CustomError
	}{
		{
			name:        "nil request",
			setup:       func() {},
			req:         nil,
			expectError: nil,
		},
		{
			name:  "invalid request - invalid mobile number",
			setup: func() {},
			req: &featuresdtos.CreateCustomerRequest{
				MobileNumber: "12345678900",
				Email:        "test@test.com",
				FirstName:    "test",
				LastName:     "test",
			},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:  "invalid request - invalid email",
			setup: func() {},
			req: &featuresdtos.CreateCustomerRequest{
				MobileNumber: "1234567890",
				Email:        "invalid",
				FirstName:    "test",
				LastName:     "test",
			},
			expectError: customizeerrors.InvalidEmailError,
		},
		{
			name:  "invalid request - invalid name",
			setup: func() {},
			req: &featuresdtos.CreateCustomerRequest{
				MobileNumber: "1234567890",
				Email:        "test@test.com",
				FirstName:    "",
				LastName:     "test",
			},
			expectError: customizeerrors.InvalidNameError,
		},
		{
			name: "successful customer creation",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockReadCustomerByMobileNumberAndActiveSwitchRepository.EXPECT().
					ReadCustomerByMobileNumberAndActiveSwitch(ctx, "1234567890", nil).
					Return(nil, nil).
					AnyTimes()
				mockCreateCustomerRepository.EXPECT().CreateCustomer(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req: &featuresdtos.CreateCustomerRequest{
				MobileNumber: "1234567890",
				Email:        "test@test.com",
				FirstName:    "test",
				LastName:     "test",
			},
			expectError: nil,
		},
		{
			name: "customer already exists",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockReadCustomerByMobileNumberAndActiveSwitchRepository.EXPECT().
					ReadCustomerByMobileNumberAndActiveSwitch(ctx, "1111111111", nil).
					Return(entities.Customers{
						entities.Customer{
							ID:           "1111111111",
							MobileNumber: "1111111111",
							Email:        "test@test.com",
							FirstName:    "test",
							LastName:     "test",
							ActiveSwitch: true,
						},
					}, nil).
					AnyTimes()
			},
			req: &featuresdtos.CreateCustomerRequest{
				MobileNumber: "1111111111",
				Email:        "test@test.com",
				FirstName:    "test",
				LastName:     "test",
			},
			expectError: customizeerrors.CustomerAlreadyExistsError,
		},
		{
			name: "customer exists but inactive",
			setup: func() {
				mockUUIDGenerator.EXPECT().GenerateUUID().Return("1234567890").AnyTimes()
				mockReadCustomerByMobileNumberAndActiveSwitchRepository.EXPECT().
					ReadCustomerByMobileNumberAndActiveSwitch(ctx, "1111111112", nil).
					Return(entities.Customers{
						entities.Customer{
							ID:           "1111111112",
							MobileNumber: "1111111112",
							Email:        "test2@test.com",
							FirstName:    "test2",
							LastName:     "test2",
							ActiveSwitch: false,
						},
					}, nil).
					AnyTimes()
			},
			req: &featuresdtos.CreateCustomerRequest{
				MobileNumber: "1111111112",
				Email:        "test2@test.com",
				FirstName:    "test2",
				LastName:     "test2",
			},
			expectError: customizeerrors.CustomerExistsButInactiveError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.CreateCustomer(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
