package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
)

type CreateAccount interface {
	CreateAccount(ctx context.Context, account dtos.Account) error
}

type GetAccountWithHistoryByMobileNumber interface {
	GetAccountWithHistoryByMobileNumber(ctx context.Context, mobileNumber string) (*dtos.AccountWithHistory, error)
}

type UpdateAccountByID interface {
	UpdateAccountByID(ctx context.Context, update dtos.UpdateAccount) error
}

type DeleteAccountByID interface {
	DeleteAccountByID(ctx context.Context, id string) error
}
