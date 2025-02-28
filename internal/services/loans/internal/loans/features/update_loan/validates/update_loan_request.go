package validates

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/dtos"
)

func ValidateUpdateLoanRequest(req featuresdtos.UpdateLoanRequest) error {
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	if req.MobileNumber != nil {
		if err := validates.ValidateMobileNumber(*req.MobileNumber); err != nil {
			return err
		}
	}
	if req.TotalAmount != nil {
		if err := validates.ValidateDecimal(*req.TotalAmount); err != nil {
			return err
		}
	}
	if req.PaidAmount != nil {
		if err := validates.ValidateDecimal(*req.PaidAmount); err != nil {
			return err
		}
	}
	if req.InterestRate != nil {
		if err := validates.ValidateDecimal(*req.InterestRate); err != nil {
			return err
		}
	}
	if req.Term != nil {
		if *req.Term <= 0 {
			return customizeerrors.LoanTermInvalidError
		}
	}
	return nil
}
