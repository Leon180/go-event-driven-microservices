package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	customErrors "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/errors"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
)

type GetAccount interface {
	GetAccount(ctx context.Context, req *featuredtos.GetAccountRequest) (*dtos.AccountWithHistory, error)
}

func NewGetAccount(
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber,
) GetAccount {
	return &getAccountImpl{getAccountWithHistoryByMobileNumberRepository: getAccountWithHistoryByMobileNumberRepository}
}

type getAccountImpl struct {
	getAccountWithHistoryByMobileNumberRepository repositories.GetAccountWithHistoryByMobileNumber
}

func (handle *getAccountImpl) GetAccount(ctx context.Context, req *featuredtos.GetAccountRequest) (*dtos.AccountWithHistory, error) {
	if req == nil {
		return nil, nil
	}
	account, err := handle.getAccountWithHistoryByMobileNumberRepository.GetAccountWithHistoryByMobileNumber(ctx, req.MobileNumber)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, customErrors.AccountNotFoundError
	}
	return account, nil
}
