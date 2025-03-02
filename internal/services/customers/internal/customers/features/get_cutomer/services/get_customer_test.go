package services

import (
	"context"
	"testing"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/dtos"
	mocksrepositories "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCustomerByMobileNumberAndActiveSwitch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReadCustomerByMobileNumberAndActiveSwitchRepository := mocksrepositories.NewMockReadCustomerByMobileNumberAndActiveSwitch(
		ctrl,
	)

	service := NewGetCustomerByMobileNumberAndActiveSwitch(
		mockReadCustomerByMobileNumberAndActiveSwitchRepository,
	)

	ctx := context.Background()

	// Test cases
	tests := []struct {
		name         string
		setup        func()
		req          *featuresdtos.GetCustomerRequest
		expectResult entities.Customers
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
			req:          &featuresdtos.GetCustomerRequest{MobileNumber: "123456789000"},
			expectResult: nil,
			expectError:  customizeerrors.InvalidMobileNumberError,
		},
		{
			name: "credit card not found",
			setup: func() {
				mockReadCustomerByMobileNumberAndActiveSwitchRepository.EXPECT().
					ReadCustomerByMobileNumberAndActiveSwitch(ctx, "9999999999", gomock.Any()).
					Return(nil, nil).
					AnyTimes()
			},
			req:          &featuresdtos.GetCustomerRequest{MobileNumber: "9999999999"},
			expectResult: nil,
			expectError:  customizeerrors.CustomerNotFoundError,
		},
		{
			name: "customer found",
			setup: func() {
				mockReadCustomerByMobileNumberAndActiveSwitchRepository.EXPECT().
					ReadCustomerByMobileNumberAndActiveSwitch(ctx, "1234567890", nil).
					Return(entities.Customers{
						entities.Customer{
							ID:           "1234567890",
							MobileNumber: "1234567890",
							Email:        "test@test.com",
							FirstName:    "Test",
							LastName:     "Test",
							ActiveSwitch: true,
						},
					}, nil).
					AnyTimes()
			},
			req: &featuresdtos.GetCustomerRequest{MobileNumber: "1234567890"},
			expectResult: entities.Customers{
				entities.Customer{
					ID:           "1234567890",
					MobileNumber: "1234567890",
					Email:        "test@test.com",
					FirstName:    "Test",
					LastName:     "Test",
					ActiveSwitch: true,
				},
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			result, err := service.GetCustomerByMobileNumberAndActiveSwitch(ctx, test.req)
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
					assert.Equal(t, account.Email, result[i].Email)
					assert.Equal(t, account.FirstName, result[i].FirstName)
					assert.Equal(t, account.LastName, result[i].LastName)
					assert.Equal(t, account.ActiveSwitch, result[i].ActiveSwitch)
				}
			}
		})
	}
}
