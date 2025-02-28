package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
	"gorm.io/gorm"
)

//go:generate mockgen -source=credit_cards_update_impl.go -destination=./mocks/credit_cards_update_impl_mock.go -package=mocks

type updateCreditCardByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewUpdateCreditCardByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateCreditCardByID {
	return &updateCreditCardByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *updateCreditCardByIDImpl) UpdateCreditCardByID(ctx context.Context, update entities.UpdateCreditCard) error {
	if update.NoUpdates() {
		return nil
	}
	updateMap := update.ToUpdateMap()
	if err := handle.db.WithContext(ctx).Model(&entities.CreditCard{}).Where("id = ?", update.ID).Updates(updateMap).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error updating credit card by id")
		return err
	}
	return nil
}
