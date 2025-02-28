package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/delete_loan/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
)

type DeleteLoan interface {
	DeleteLoan(ctx context.Context, req *featuresdtos.DeleteLoanRequest) error
}

func NewDeleteLoan(
	readLoanRepository repositories.ReadLoan,
	updateLoanByIDRepository repositories.UpdateLoanByID,
) DeleteLoan {
	return &deleteLoanImpl{
		readLoanRepository:       readLoanRepository,
		updateLoanByIDRepository: updateLoanByIDRepository,
	}
}

type deleteLoanImpl struct {
	readLoanRepository       repositories.ReadLoan
	updateLoanByIDRepository repositories.UpdateLoanByID
}

func (handle *deleteLoanImpl) DeleteLoan(ctx context.Context, req *featuresdtos.DeleteLoanRequest) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	loan, err := handle.readLoanRepository.ReadLoan(ctx, req.ID)
	if err != nil {
		return err
	}
	if loan == nil {
		return customizeerrors.LoanNotFoundError
	}
	if !loan.IsActive() {
		return customizeerrors.LoanAlreadyDeletedError
	}
	activeSwitch := false
	updateLoan := entities.UpdateLoan{
		ID:           loan.ID,
		ActiveSwitch: &activeSwitch,
	}
	return handle.updateLoanByIDRepository.UpdateLoanByID(ctx, updateLoan)
}
