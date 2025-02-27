package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
)

type UpdateAccountByID interface {
	UpdateAccountByID(ctx context.Context, update dtos.UpdateAccount) error
}
