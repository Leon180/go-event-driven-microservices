package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
)

type GetLoansByMobileNumberAndActiveSwitch interface {
	GetLoansByMobileNumberAndActiveSwitch(ctx context.Context, req *featuresdtos.GetLoansRequest) (entities.Loans, error)
}

func NewGetLoansByMobileNumberAndActiveSwitch(
	readLoanByMobileNumberAndActiveSwitchRepository repositories.ReadLoanByMobileNumberAndActiveSwitch,
) GetLoansByMobileNumberAndActiveSwitch {
	return &getLoansByMobileNumberAndActiveSwitchImpl{readLoanByMobileNumberAndActiveSwitchRepository: readLoanByMobileNumberAndActiveSwitchRepository}
}

type getLoansByMobileNumberAndActiveSwitchImpl struct {
	readLoanByMobileNumberAndActiveSwitchRepository repositories.ReadLoanByMobileNumberAndActiveSwitch
}

func (handle *getLoansByMobileNumberAndActiveSwitchImpl) GetLoansByMobileNumberAndActiveSwitch(ctx context.Context, req *featuresdtos.GetLoansRequest) (entities.Loans, error) {
	if req == nil {
		return nil, nil
	}
	if err := validates.ValidateGetLoansRequest(req); err != nil {
		return nil, err
	}
	loans, err := handle.readLoanByMobileNumberAndActiveSwitchRepository.ReadLoanByMobileNumberAndActiveSwitch(ctx, req.MobileNumber, req.ActiveSwitch)
	if err != nil {
		return nil, err
	}
	if len(loans) == 0 {
		return nil, customizeerrors.LoanNotFoundError
	}
	return loans, nil
}
