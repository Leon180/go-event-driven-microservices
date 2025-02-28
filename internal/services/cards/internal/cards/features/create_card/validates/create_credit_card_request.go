package validates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/dtos"
)

func ValidateCreateCreditCardRequest(req *featuresdtos.CreateCreditCardRequest) error {
	if err := validates.ValidateMobileNumber(req.MobileNumber); err != nil {
		return err
	}
	if err := validates.ValidateDecimal(req.TotalLimit); err != nil {
		return err
	}
	return nil
}
