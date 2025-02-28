package validates

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/dtos"
)

func ValidateUpdateCreditCardRequest(req featuresdtos.UpdateCreditCardRequest) error {
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	if req.MobileNumber != nil {
		if err := validates.ValidateMobileNumber(*req.MobileNumber); err != nil {
			return err
		}
	}
	if req.TotalLimit != nil {
		if err := validates.ValidateDecimal(*req.TotalLimit); err != nil {
			return err
		}
	}
	if req.AmountUsed != nil {
		if err := validates.ValidateDecimal(*req.AmountUsed); err != nil {
			return err
		}
	}
	return nil
}
