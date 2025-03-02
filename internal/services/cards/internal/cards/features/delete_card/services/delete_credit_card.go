package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/delete_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
)

type DeleteCreditCard interface {
	DeleteCreditCard(ctx context.Context, req *featuresdtos.DeleteCreditCardRequest) error
}

func NewDeleteCreditCard(
	readCreditCardRepository repositories.ReadCreditCard,
	updateCreditCardByIDRepository repositories.UpdateCreditCardByID,
) DeleteCreditCard {
	return &deleteCreditCardImpl{
		readCreditCardRepository:       readCreditCardRepository,
		updateCreditCardByIDRepository: updateCreditCardByIDRepository,
	}
}

type deleteCreditCardImpl struct {
	readCreditCardRepository       repositories.ReadCreditCard
	updateCreditCardByIDRepository repositories.UpdateCreditCardByID
}

func (handle *deleteCreditCardImpl) DeleteCreditCard(
	ctx context.Context,
	req *featuresdtos.DeleteCreditCardRequest,
) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	card, err := handle.readCreditCardRepository.ReadCreditCard(ctx, req.ID)
	if err != nil {
		return err
	}
	if card == nil {
		return customizeerrors.CardNotFoundError
	}
	if !card.IsActive() {
		return customizeerrors.CardAlreadyDeletedError
	}
	activeSwitch := false
	updateCard := entities.UpdateCreditCard{
		ID:           card.ID,
		ActiveSwitch: &activeSwitch,
	}
	return handle.updateCreditCardByIDRepository.UpdateCreditCardByID(ctx, updateCard)
}
