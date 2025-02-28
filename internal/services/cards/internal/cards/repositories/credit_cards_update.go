package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
)

//go:generate mockgen -source=credit_cards_update.go -destination=./mocks/credit_cards_update_mock.go -package=mocks

type UpdateCreditCardByID interface {
	UpdateCreditCardByID(ctx context.Context, update entities.UpdateCreditCard) error
}
