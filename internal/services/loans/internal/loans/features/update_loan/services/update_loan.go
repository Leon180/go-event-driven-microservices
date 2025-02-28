package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
	"github.com/shopspring/decimal"
)

type UpdateLoan interface {
	UpdateLoan(ctx context.Context, req *featuresdtos.UpdateLoanRequest) error
}

type updateLoanImpl struct {
	readLoan                 repositories.ReadLoan
	updateLoanByIDRepository repositories.UpdateLoanByID
}

func NewUpdateLoan(
	readLoan repositories.ReadLoan,
	updateLoanByIDRepository repositories.UpdateLoanByID,
) UpdateLoan {
	return &updateLoanImpl{
		readLoan:                 readLoan,
		updateLoanByIDRepository: updateLoanByIDRepository,
	}
}

func (handle *updateLoanImpl) UpdateLoan(ctx context.Context, req *featuresdtos.UpdateLoanRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateUpdateLoanRequest(*req); err != nil {
		return err
	}
	loan, err := handle.readLoan.ReadLoan(ctx, req.ID)
	if err != nil {
		return err
	}
	if loan == nil {
		return customizeerrors.LoanNotFoundError
	}
	updateLoan := entities.UpdateLoan{
		ID:           req.ID,
		MobileNumber: req.MobileNumber,
		TotalAmount: func() *decimal.Decimal {
			if req.TotalAmount != nil {
				totalAmount, _ := decimal.NewFromString(*req.TotalAmount)
				return &totalAmount
			}
			return nil
		}(),
		// add amount paid to the loan
		PaidAmount: func() *decimal.Decimal {
			if req.PaidAmount != nil {
				paidAmount, _ := decimal.NewFromString(*req.PaidAmount)
				paidAmount = paidAmount.Add(loan.PaidAmount)
				return &paidAmount
			}
			return nil
		}(),
		InterestRate: func() *decimal.Decimal {
			if req.InterestRate != nil {
				interestRate, _ := decimal.NewFromString(*req.InterestRate)
				return &interestRate
			}
			return nil
		}(),
		// add term to the loan
		Term: func() *int {
			if req.Term != nil {
				t := *req.Term + loan.Term
				return &t
			}
			return nil
		}(),
		ActiveSwitch: req.ActiveSwitch,
	}
	updateLoan.RemoveUnchangedFields(*loan)
	if updateLoan.NoUpdates() {
		return customizeerrors.LoanNoUpdatesError
	}
	return handle.updateLoanByIDRepository.UpdateLoanByID(ctx, updateLoan)
}
