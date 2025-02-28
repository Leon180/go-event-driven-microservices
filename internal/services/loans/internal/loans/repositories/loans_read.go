package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
)

//go:generate mockgen -source=loans_read.go -destination=mocks/loans_read_mock.go -package=mocks

type ReadLoanByMobileNumberAndActiveSwitch interface {
	ReadLoanByMobileNumberAndActiveSwitch(ctx context.Context, mobileNumber string, activeSwitch *bool) (entities.Loans, error)
}

type ReadLoan interface {
	ReadLoan(ctx context.Context, id string) (*entities.Loan, error)
}
