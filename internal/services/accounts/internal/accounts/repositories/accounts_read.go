package repositories

import (
	"context"

	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
)

type ReadAccountsByMobileNumber interface {
	ReadAccountsByMobileNumber(ctx context.Context, mobileNumber string) (entities.Accounts, error)
}

type ReadAccountByMobileNumberAndAccountType interface {
	ReadAccountByMobileNumberAndAccountType(ctx context.Context, mobileNumber string, accountTypeCode enumsaccounts.AccountTypeCode) (*entities.Account, error)
}

type ReadAccount interface {
	ReadAccount(ctx context.Context, id string) (*entities.Account, error)
}
