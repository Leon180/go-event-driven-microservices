package repositories

import (
	"context"

	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
)

type ReadAccountsWithHistoryByMobileNumber interface {
	ReadAccountsWithHistoryByMobileNumber(ctx context.Context, mobileNumber string) ([]dtos.AccountWithHistory, error)
}

type ReadAccountWithHistoryByMobileNumberAndAccountType interface {
	ReadAccountWithHistoryByMobileNumberAndAccountType(ctx context.Context, mobileNumber string, accountTypeCode enumsaccounts.AccountTypeCode) (*dtos.AccountWithHistory, error)
}

type ReadAccountWithHistory interface {
	ReadAccountWithHistory(ctx context.Context, id string) (*dtos.AccountWithHistory, error)
}
