package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
)

type UpdateAccountByID interface {
	UpdateAccountByID(ctx context.Context, update entities.UpdateAccount) error
}
