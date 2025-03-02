package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
)

type GetCreditCardsByMobileNumberAndActiveSwitch interface {
	GetCreditCardsByMobileNumberAndActiveSwitch(
		ctx context.Context,
		req *featuresdtos.GetCreditCardsRequest,
	) (entities.CreditCards, error)
}

func NewGetCreditCardsByMobileNumberAndActiveSwitch(
	readCreditCardByMobileNumberAndActiveSwitchRepository repositories.ReadCreditCardByMobileNumberAndActiveSwitch,
) GetCreditCardsByMobileNumberAndActiveSwitch {
	return &getCreditCardsByMobileNumberAndActiveSwitchImpl{
		readCreditCardByMobileNumberAndActiveSwitchRepository: readCreditCardByMobileNumberAndActiveSwitchRepository,
	}
}

type getCreditCardsByMobileNumberAndActiveSwitchImpl struct {
	readCreditCardByMobileNumberAndActiveSwitchRepository repositories.ReadCreditCardByMobileNumberAndActiveSwitch
}

func (handle *getCreditCardsByMobileNumberAndActiveSwitchImpl) GetCreditCardsByMobileNumberAndActiveSwitch(
	ctx context.Context,
	req *featuresdtos.GetCreditCardsRequest,
) (entities.CreditCards, error) {
	if req == nil {
		return nil, nil
	}
	if err := validates.ValidateGetCreditCardsRequest(req); err != nil {
		return nil, err
	}
	creditCards, err := handle.readCreditCardByMobileNumberAndActiveSwitchRepository.ReadCreditCardByMobileNumberAndActiveSwitch(
		ctx,
		req.MobileNumber,
		req.ActiveSwitch,
	)
	if err != nil {
		return nil, err
	}
	if len(creditCards) == 0 {
		return nil, customizeerrors.CardNotFoundError
	}
	return creditCards, nil
}
