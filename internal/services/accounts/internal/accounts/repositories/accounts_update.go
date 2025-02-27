package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
)

//go:generate mockgen -source=accounts_update.go -destination=./mocks/accounts_update_mock.go -package=mocks

type UpdateAccountByID interface {
	UpdateAccountByID(ctx context.Context, update entities.UpdateAccount) error
}
