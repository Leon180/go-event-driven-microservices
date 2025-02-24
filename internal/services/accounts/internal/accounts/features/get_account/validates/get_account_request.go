package validates

import (
	"net/http"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/dtos"
)

func ValidateGetAccountRequest(req featuresdtos.GetAccountRequest) error {
	if req.MobileNumber == "" {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Account number is required")
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Mobile number is invalid")
	}
	return nil
}
