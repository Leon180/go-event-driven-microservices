package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
)

type CreateAccount interface {
	CreateAccount(ctx context.Context, account entities.Account) error
}

type CreateAccounts interface {
	CreateAccounts(ctx context.Context, accounts entities.Accounts) error
}
