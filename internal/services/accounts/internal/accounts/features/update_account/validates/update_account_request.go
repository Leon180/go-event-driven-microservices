package validates

import (
	"net/http"
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
)

func ValidateUpdateAccountRequest(req featuresdtos.UpdateAccountRequest) error {
	if req.MobileNumber == "" {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Account number is required")
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Mobile number is invalid")
	}
	if req.AccountType != nil && !enumsaccounts.AccountType(strings.ToLower(*req.AccountType)).IsValid() {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Account type is invalid")
	}
	if req.BranchAddress != nil && !enumsbanks.BanksBranch(strings.ToLower(*req.BranchAddress)).IsValid() {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "Branch address is invalid")
	}
	return nil
}
