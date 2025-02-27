package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type GetAccountsByMobileNumber interface {
	GetAccountsByMobileNumber(ctx context.Context, req *featuresdtos.GetAccountsByMobileNumberRequest) (entities.Accounts, error)
}

func NewGetAccountsByMobileNumber(
	readAccountsByMobileNumberRepository repositories.ReadAccountsByMobileNumber,
) GetAccountsByMobileNumber {
	return &getAccountsByMobileNumberImpl{readAccountsByMobileNumberRepository: readAccountsByMobileNumberRepository}
}

type getAccountsByMobileNumberImpl struct {
	readAccountsByMobileNumberRepository repositories.ReadAccountsByMobileNumber
}

func (handle *getAccountsByMobileNumberImpl) GetAccountsByMobileNumber(ctx context.Context, req *featuresdtos.GetAccountsByMobileNumberRequest) (entities.Accounts, error) {
	if req == nil {
		return nil, nil
	}
	if err := validates.ValidateGetAccountsByMobileNumberRequest(*req); err != nil {
		return nil, err
	}
	accounts, err := handle.readAccountsByMobileNumberRepository.ReadAccountsByMobileNumber(ctx, req.MobileNumber)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, customizeerrors.AccountNotFoundError
	}
	return accounts, nil
}
