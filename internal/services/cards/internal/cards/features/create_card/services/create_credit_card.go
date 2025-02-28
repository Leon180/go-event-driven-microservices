package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
	cardnumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/utilities/card_number"
	"github.com/shopspring/decimal"
)

type CreateCreditCard interface {
	CreateCreditCard(ctx context.Context, req *featuresdtos.CreateCreditCardRequest) error
}

func NewCreateCreditCard(
	createCreditCardRepository repositories.CreateCreditCard,
	readCreditCardByMobileNumberAndActiveSwitch repositories.ReadCreditCardByMobileNumberAndActiveSwitch,
	uuidGenerator uuid.UUIDGenerator,
	cardNumberGenerator cardnumberutilities.CardNumberGenerator,
) CreateCreditCard {
	return &createCreditCardImpl{
		createCreditCardRepository:                  createCreditCardRepository,
		readCreditCardByMobileNumberAndActiveSwitch: readCreditCardByMobileNumberAndActiveSwitch,
		uuidGenerator:       uuidGenerator,
		cardNumberGenerator: cardNumberGenerator,
	}
}

type createCreditCardImpl struct {
	createCreditCardRepository                  repositories.CreateCreditCard
	readCreditCardByMobileNumberAndActiveSwitch repositories.ReadCreditCardByMobileNumberAndActiveSwitch
	uuidGenerator                               uuid.UUIDGenerator
	cardNumberGenerator                         cardnumberutilities.CardNumberGenerator
}

func (handle *createCreditCardImpl) CreateCreditCard(ctx context.Context, req *featuresdtos.CreateCreditCardRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateCreateCreditCardRequest(req); err != nil {
		return err
	}

	activeSwitch := true
	cards, err := handle.readCreditCardByMobileNumberAndActiveSwitch.ReadCreditCardByMobileNumberAndActiveSwitch(ctx, req.MobileNumber, &activeSwitch)
	if err != nil {
		return err
	}
	if len(cards) > 0 {
		return customizeerrors.CardAlreadyExistsError
	}

	systemTime := time.Now()
	cardEntity := entities.CreditCard{
		ID:           handle.uuidGenerator.GenerateUUID(),
		MobileNumber: req.MobileNumber,
		CardNumber:   handle.cardNumberGenerator.GenerateCardNumber(),
		TotalLimit: func() decimal.Decimal {
			totalLimit, _ := decimal.NewFromString(req.TotalLimit)
			return totalLimit
		}(),
		AmountUsed:   decimal.Zero,
		ActiveSwitch: true,
		CommonHistoryModelWithUpdate: entities.CommonHistoryModelWithUpdate{
			CommonHistoryModel: entities.CommonHistoryModel{
				CreatedAt: systemTime,
			},
			UpdatedAt: systemTime,
		},
	}

	return handle.createCreditCardRepository.CreateCreditCard(ctx, &cardEntity)
}
