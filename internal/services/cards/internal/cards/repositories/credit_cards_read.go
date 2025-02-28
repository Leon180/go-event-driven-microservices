package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
)

//go:generate mockgen -source=credit_cards_read.go -destination=mocks/credit_cards_read_mock.go -package=mocks

type ReadCreditCardByMobileNumberAndActiveSwitch interface {
	ReadCreditCardByMobileNumberAndActiveSwitch(ctx context.Context, mobileNumber string, activeSwitch *bool) (entities.CreditCards, error)
}

type ReadCreditCard interface {
	ReadCreditCard(ctx context.Context, id string) (*entities.CreditCard, error)
}
