package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadCustomerRepository := mocksrepositories.NewMockReadCustomer(ctrl)
	mockUpdateCustomerByIDRepository := mocksrepositories.NewMockUpdateCustomerByID(ctrl)

	service := NewUpdateCustomer(
		mockReadCustomerRepository,
		mockUpdateCustomerByIDRepository,
	)

	ctx := context.Background()

	invalidMobileNumber := "123456789000"
	validMobileNumber := "1234567890"
	updateMobileNumber := "1234567891"
	toMobileNumber := "1234567830"
	invalidEmail := "invalid"
	validEmail := "test@test.com"
	invalidFirstName := ""
	validFirstName := "John"
	validLastName := "Doe"

	// Test cases
	tests := []struct {
		name        string
		setup       func()
		req         *featuresdtos.UpdateCustomerRequest
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
			req:         &featuresdtos.UpdateCustomerRequest{ID: ""},
			expectError: customizeerrors.InvalidIDError,
		},
		{
			name:        "invalid request - invalid mobile number",
			setup:       func() {},
			req:         &featuresdtos.UpdateCustomerRequest{ID: "1234567890", MobileNumber: &invalidMobileNumber},
			expectError: customizeerrors.InvalidMobileNumberError,
		},
		{
			name:        "invalid request - invalid email",
			setup:       func() {},
			req:         &featuresdtos.UpdateCustomerRequest{ID: "1234567890", Email: &invalidEmail},
			expectError: customizeerrors.InvalidEmailError,
		},
		{
			name:        "invalid request - invalid name",
			setup:       func() {},
			req:         &featuresdtos.UpdateCustomerRequest{ID: "1234567890", FirstName: &invalidFirstName},
			expectError: customizeerrors.InvalidNameError,
		},
		{
			name: "customer not found",
			setup: func() {
				mockReadCustomerRepository.EXPECT().ReadCustomer(ctx, "123456789000").Return(nil, nil).AnyTimes()
			},
			req:         &featuresdtos.UpdateCustomerRequest{ID: "123456789000"},
			expectError: customizeerrors.CustomerNotFoundError,
		},
		{
			name: "customer no updates",
			setup: func() {
				mockReadCustomerRepository.EXPECT().ReadCustomer(ctx, "1111111111").Return(&entities.Customer{
					ID:           "1111111111",
					MobileNumber: validMobileNumber,
					Email:        validEmail,
					FirstName:    validFirstName,
					LastName:     validLastName,
					ActiveSwitch: true,
				}, nil).AnyTimes()
			},
			req: &featuresdtos.UpdateCustomerRequest{
				ID:           "1111111111",
				MobileNumber: &validMobileNumber,
				Email:        &validEmail,
				FirstName:    &validFirstName,
				LastName:     &validLastName,
			},
			expectError: customizeerrors.CustomerNoUpdatesError,
		},
		{
			name: "customer successfully updated",
			setup: func() {
				mockReadCustomerRepository.EXPECT().ReadCustomer(ctx, "1234567890").Return(&entities.Customer{
					ID:           "1234567890",
					MobileNumber: updateMobileNumber,
					Email:        validEmail,
					FirstName:    validFirstName,
					LastName:     validLastName,
					ActiveSwitch: true,
				}, nil).AnyTimes()
				mockUpdateCustomerByIDRepository.EXPECT().UpdateCustomerByID(ctx, gomock.Any()).Return(nil).AnyTimes()
			},
			req:         &featuresdtos.UpdateCustomerRequest{ID: "1234567890", MobileNumber: &toMobileNumber},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			err := service.UpdateCustomer(ctx, test.req)
			if test.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expectError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
