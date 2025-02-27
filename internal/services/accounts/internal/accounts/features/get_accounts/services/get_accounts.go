package services

import (
	"context"

	customizerrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type GetAccountsByMobileNumber interface {
	GetAccountsByMobileNumber(ctx context.Context, req *featuresdtos.GetAccountsByMobileNumberRequest) ([]dtos.AccountWithHistory, error)
}

func NewGetAccountsByMobileNumber(
	readAccountsByMobileNumberRepository repositories.ReadAccountsWithHistoryByMobileNumber,
) GetAccountsByMobileNumber {
	return &getAccountsByMobileNumberImpl{readAccountsByMobileNumberRepository: readAccountsByMobileNumberRepository}
}

type getAccountsByMobileNumberImpl struct {
	readAccountsByMobileNumberRepository repositories.ReadAccountsWithHistoryByMobileNumber
}

func (handle *getAccountsByMobileNumberImpl) GetAccountsByMobileNumber(ctx context.Context, req *featuresdtos.GetAccountsByMobileNumberRequest) ([]dtos.AccountWithHistory, error) {
	if req == nil {
		return nil, nil
	}
	if err := validates.ValidateGetAccountsByMobileNumberRequest(*req); err != nil {
		return nil, err
	}
	accounts, err := handle.readAccountsByMobileNumberRepository.ReadAccountsWithHistoryByMobileNumber(ctx, req.MobileNumber)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, customizerrors.AccountNotFoundError
	}
	return accounts, nil
}
