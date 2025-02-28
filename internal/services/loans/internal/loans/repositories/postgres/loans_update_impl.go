package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
	"gorm.io/gorm"
)

type updateLoanByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewUpdateLoanByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateLoanByID {
	return &updateLoanByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *updateLoanByIDImpl) UpdateLoanByID(ctx context.Context, update entities.UpdateLoan) error {
	if update.NoUpdates() {
		return nil
	}
	updateMap := update.ToUpdateMap()
	if err := handle.db.WithContext(ctx).Model(&entities.Loan{}).Where("id = ?", update.ID).Updates(updateMap).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error updating loan by id")
		return err
	}
	return nil
}
