package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
)

type CreateAccount interface {
	CreateAccount(ctx context.Context, account dtos.Account) error
}

type CreateAccounts interface {
	CreateAccounts(ctx context.Context, accounts []dtos.Account) error
}
