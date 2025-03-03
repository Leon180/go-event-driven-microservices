package validates

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/dtos"
)

func ValidateGetAccountsByMobileNumberRequest(req featuresdtos.GetAccountsByMobileNumberRequest) error {
	if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.InvalidMobileNumberError
	}
	return nil
}
