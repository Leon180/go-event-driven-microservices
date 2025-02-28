package validates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/dtos"
)

func ValidateGetCreditCardsRequest(req *featuresdtos.GetCreditCardsRequest) error {
	return validates.ValidateMobileNumber(req.MobileNumber)
}
