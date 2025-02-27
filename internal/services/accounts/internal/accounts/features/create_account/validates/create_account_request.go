package validates

import (
	"net/http"
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
)

func ValidateCreateAccountRequest(req *featuresdtos.CreateAccountRequest) error {
	if !enumsaccounts.AccountType(strings.ToLower(req.AccountType)).IsValid() {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "account type is invalid")
	}
	if !enumsbanks.BanksBranch(strings.ToLower(req.Branch)).IsValid() {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "branch is invalid")
	}
	if len(req.MobileNumber) == 0 {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "mobile number is required")
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.NewError(http.StatusBadRequest, int(customizeerrors.HTTPBadRequest), "mobile number is invalid")
	}
	return nil
}
