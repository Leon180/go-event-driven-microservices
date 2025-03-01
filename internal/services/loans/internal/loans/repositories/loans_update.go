package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
)

//go:generate mockgen -source=loans_update.go -destination=./mocks/loans_update_mock.go -package=mocks

type UpdateLoanByID interface {
	UpdateLoanByID(ctx context.Context, update entities.UpdateLoan) error
}
