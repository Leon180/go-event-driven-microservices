package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
)

//go:generate mockgen -source=loans_create.go -destination=mocks/loans_create_mock.go -package=mocks

type CreateLoan interface {
	CreateLoan(ctx context.Context, loan *entities.Loan) error
}

type CreateLoans interface {
	CreateLoans(ctx context.Context, loans entities.Loans) error
}
