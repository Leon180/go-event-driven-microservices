package validates

import (
	"net/http"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customErrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/dtos"
)

func ValidateGetAccountRequest(req dtos.GetAccountRequest) error {
	if req.MobileNumber == "" {
		return customErrors.NewError(http.StatusBadRequest, int(customErrors.HTTPBadRequest), "Account number is required")
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customErrors.NewError(http.StatusBadRequest, int(customErrors.HTTPBadRequest), "Mobile number is invalid")
	}
	return nil
}
