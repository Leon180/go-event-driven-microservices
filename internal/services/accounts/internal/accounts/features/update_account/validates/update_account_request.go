package validates

import (
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	enumsbanks "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/banks"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
)

func ValidateUpdateAccountRequest(req featuresdtos.UpdateAccountRequest) error {
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	if req.MobileNumber != nil {
		if !enums.MobileNumberFormat.ValidateFormat(*req.MobileNumber) {
			return customizeerrors.InvalidMobileNumberError
		}
	}
	if req.BranchAddress != nil && !enumsbanks.BanksBranch(strings.ToLower(*req.BranchAddress)).IsValid() {
		return customizeerrors.InvalidBranchError
	}
	return nil
}
