package validates

import (
	"net/http"
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
)

func ValidateUpdateAccountRequest(req featuresdtos.UpdateAccountRequest) error {
	if req.MobileNumber != nil {
		if *req.MobileNumber == "" {
			return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Account number is required")
		} else if !enums.MobileNumberFormat.ValidateFormat(*req.MobileNumber) {
			return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Mobile number is invalid")
		}
	}
	if req.BranchAddress != nil && !enumsbanks.BanksBranch(strings.ToLower(*req.BranchAddress)).IsValid() {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Branch address is invalid")
	}
	return nil
}
