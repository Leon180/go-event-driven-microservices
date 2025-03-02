package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
	"github.com/shopspring/decimal"
)

type UpdateCreditCard interface {
	UpdateCreditCard(ctx context.Context, req *featuresdtos.UpdateCreditCardRequest) error
}

type updateCreditCardImpl struct {
	readCreditCard                 repositories.ReadCreditCard
	updateCreditCardByIDRepository repositories.UpdateCreditCardByID
}

func NewUpdateCreditCard(
	readCreditCard repositories.ReadCreditCard,
	updateCreditCardByIDRepository repositories.UpdateCreditCardByID,
) UpdateCreditCard {
	return &updateCreditCardImpl{
		readCreditCard:                 readCreditCard,
		updateCreditCardByIDRepository: updateCreditCardByIDRepository,
	}
}

func (handle *updateCreditCardImpl) UpdateCreditCard(
	ctx context.Context,
	req *featuresdtos.UpdateCreditCardRequest,
) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateUpdateCreditCardRequest(*req); err != nil {
		return err
	}
	creditCard, err := handle.readCreditCard.ReadCreditCard(ctx, req.ID)
	if err != nil {
		return err
	}
	if creditCard == nil {
		return customizeerrors.CardNotFoundError
	}
	updateCreditCard := entities.UpdateCreditCard{
		ID:           req.ID,
		MobileNumber: req.MobileNumber,
		TotalLimit: func() *decimal.Decimal {
			if req.TotalLimit != nil {
				totalLimit, _ := decimal.NewFromString(*req.TotalLimit)
				return &totalLimit
			}
			return nil
		}(),
		// add amount used to the credit card
		AmountUsed: func() *decimal.Decimal {
			if req.AmountUsed != nil {
				amountUsed, _ := decimal.NewFromString(*req.AmountUsed)
				amountUsed = amountUsed.Add(creditCard.AmountUsed)
				return &amountUsed
			}
			return nil
		}(),
	}
	updateCreditCard.RemoveUnchangedFields(*creditCard)
	if updateCreditCard.NoUpdates() {
		return customizeerrors.CardNoUpdatesError
	}
	return handle.updateCreditCardByIDRepository.UpdateCreditCardByID(ctx, updateCreditCard)
}
