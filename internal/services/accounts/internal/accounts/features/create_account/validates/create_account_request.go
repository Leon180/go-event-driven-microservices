package validates

import (
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
)

func ValidateCreateAccountRequest(req *featuresdtos.CreateAccountRequest) error {
	if !enumsaccounts.AccountType(strings.ToLower(req.AccountType)).IsValid() {
		return customizeerrors.InvalidAccountTypeError
	}
	if !enumsbanks.BanksBranch(strings.ToLower(req.Branch)).IsValid() {
		return customizeerrors.InvalidBranchError
	}
	if len(req.MobileNumber) == 0 {
		return customizeerrors.InvalidMobileNumberError
	} else if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.InvalidMobileNumberError
	}
	return nil
}
