package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
	"gorm.io/gorm"
)

type updateCustomerByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewUpdateCustomerByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateCustomerByID {
	return &updateCustomerByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *updateCustomerByIDImpl) UpdateCustomerByID(ctx context.Context, update entities.UpdateCustomer) error {
	if update.NoUpdates() {
		return nil
	}
	updateMap := update.ToUpdateMap()
	if err := handle.db.WithContext(ctx).Model(&entities.Customer{}).Where("id = ?", update.ID).Updates(updateMap).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error updating customer by id")
		return err
	}
	return nil
}
