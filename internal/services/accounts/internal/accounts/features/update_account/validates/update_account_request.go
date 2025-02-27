package validates

import (
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
)

func ValidateUpdateAccountRequest(req featuresdtos.UpdateAccountRequest) error {
	if req.MobileNumber != nil {
		if *req.MobileNumber == "" {
			return customizeerrors.AccountInvalidMobileNumberError
		} else if !enums.MobileNumberFormat.ValidateFormat(*req.MobileNumber) {
			return customizeerrors.AccountInvalidMobileNumberError
		}
	}
	if req.BranchAddress != nil && !enumsbanks.BanksBranch(strings.ToLower(*req.BranchAddress)).IsValid() {
		return customizeerrors.AccountInvalidBranchError
	}
	return nil
}
