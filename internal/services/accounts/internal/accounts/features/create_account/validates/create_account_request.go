package validates

import (
	"net/http"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
)

func ValidateCreateAccountRequest(req *dtos.CreateAccountRequest) error {
	if req.AccountNumber == 0 {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "account number is required")
	}
	if len(req.AccountType) == 0 {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "account type is required")
	}
	if len(req.BranchAddress) == 0 {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "branch address is required")
	}
	if len(req.MobileNumber) == 0 {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "mobile number is required")
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "mobile number is invalid")
	}
	return nil
}
