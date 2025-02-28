package services

// func TestDeleteCreditCard(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockReadCreditCardRepository := mocksrepositories.NewMockReadCreditCard(ctrl)
// 	mockUpdateCreditCardByIDRepository := mocksrepositories.NewMockUpdateCreditCardByID(ctrl)

// 	service := NewDeleteCreditCard(
// 		mockReadCreditCardRepository,
// 		mockUpdateCreditCardByIDRepository,
// 	)

// 	ctx := context.Background()

// 	// Test cases
// 	var tests = []struct {
// 		name        string
// 		setup       func()
// 		req         *featuresdtos.DeleteCreditCardRequest
// 		expectError customizeerrors.CustomError
// 	}{
// 		{
// 			name:        "nil request",
// 			setup:       func() {},
// 			req:         nil,
// 			expectError: nil,
// 		},
// 		{
// 			name:        "invalid request - empty id",
// 			setup:       func() {},
// 			req:         &featuresdtos.DeleteCreditCardRequest{ID: ""},
// 			expectError: customizeerrors.InvalidIDError,
// 		},
// 		{
// 			name: "credit card not found",
// 			setup: func() {
// 				mockReadCreditCardRepository.EXPECT().ReadCreditCard(ctx, "123456789000").Return(nil, nil).AnyTimes()
// 			},
// 			req:         &featuresdtos.DeleteCreditCardRequest{ID: "123456789000"},
// 			expectError: customizeerrors.CardNotFoundError,
// 		},
// 		{
// 			name: "credit card already deleted",
// 			setup: func() {
// 				mockReadCreditCardRepository.EXPECT().ReadCreditCard(ctx, "1111111111").Return(&entities.CreditCard{
// 					ID:           "1111111111",
// 					MobileNumber: "1111111111",
// 					CardNumber:   "1111111111",
// 					TotalLimit:   decimal.NewFromInt(100000),
// 					ActiveSwitch: false,
// 				}, nil).AnyTimes()
// 			},
// 			req:         &featuresdtos.DeleteCreditCardRequest{ID: "1111111111"},
// 			expectError: customizeerrors.CardAlreadyDeletedError,
// 		},
// 		{
// 			name: "credit card successfully deleted",
// 			setup: func() {
// 				mockReadCreditCardRepository.EXPECT().ReadCreditCard(ctx, "1234567890").Return(&entities.CreditCard{
// 					ID:           "1234567890",
// 					MobileNumber: "1234567890",
// 					CardNumber:   "1234567890",
// 					TotalLimit:   decimal.NewFromInt(100000),
// 					ActiveSwitch: true,
// 				}, nil).AnyTimes()
// 				mockUpdateCreditCardByIDRepository.EXPECT().UpdateCreditCardByID(ctx, gomock.Any()).Return(nil).AnyTimes()
// 			},
// 			req:         &featuresdtos.DeleteCreditCardRequest{ID: "1234567890"},
// 			expectError: nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup()
// 			err := service.DeleteCreditCard(ctx, test.req)
// 			if test.expectError != nil {
// 				assert.Error(t, err)
// 				assert.Equal(t, test.expectError, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
