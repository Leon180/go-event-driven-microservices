package validates

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/dtos"
)

func ValidateCreateLoanRequest(req *featuresdtos.CreateLoanRequest) error {
	if err := validates.ValidateMobileNumber(req.MobileNumber); err != nil {
		return err
	}
	if err := validates.ValidateDecimal(req.TotalAmount); err != nil {
		return err
	}
	if err := validates.ValidateDecimal(req.InterestRate); err != nil {
		return err
	}
	if req.Term <= 0 {
		return customizeerrors.LoanTermInvalidError
	}
	if !enums.LoanType(req.LoanType).IsValid() {
		return customizeerrors.InvalidLoanTypeError
	}
	return nil
}
