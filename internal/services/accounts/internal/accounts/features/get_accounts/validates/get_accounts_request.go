package validates

import (
	"net/http"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/dtos"
)

func ValidateGetAccountsByMobileNumberRequest(req featuresdtos.GetAccountsByMobileNumberRequest) error {
	if req.MobileNumber == "" {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Mobile number is required")
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Mobile number is invalid")
	}
	return nil
}
