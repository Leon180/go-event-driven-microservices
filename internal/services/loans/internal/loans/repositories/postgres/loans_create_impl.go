package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
	"gorm.io/gorm"
)

type createLoanImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateLoan(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateLoan {
	return &createLoanImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createLoanImpl) CreateLoan(ctx context.Context, loan *entities.Loan) error {
	if loan == nil {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(loan).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating loan")
		return err
	}
	return nil
}

type createLoansImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateLoans(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateLoans {
	return &createLoansImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createLoansImpl) CreateLoans(ctx context.Context, loans entities.Loans) error {
	if loans == nil {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(loans).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating loans")
		return err
	}
	return nil
}
