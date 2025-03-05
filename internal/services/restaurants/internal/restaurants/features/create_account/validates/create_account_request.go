package validates

import (
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
)

func ValidateCreateAccountRequest(req *featuresdtos.CreateAccountRequest) error {
	if !enums.AccountType(strings.ToLower(req.AccountType)).IsValid() {
		return customizeerrors.InvalidAccountTypeError
	}
	if !enums.BanksBranch(strings.ToLower(req.Branch)).IsValid() {
		return customizeerrors.InvalidBranchError
	}
	if err := validates.ValidateMobileNumber(req.MobileNumber); err != nil {
		return err
	}
	return nil
}
