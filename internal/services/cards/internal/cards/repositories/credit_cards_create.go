package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
)

//go:generate mockgen -source=credit_cards_create.go -destination=mocks/credit_cards_create_mock.go -package=mocks

type CreateCreditCard interface {
	CreateCreditCard(ctx context.Context, card *entities.CreditCard) error
}

type CreateCreditCards interface {
	CreateCreditCards(ctx context.Context, cards entities.CreditCards) error
}
