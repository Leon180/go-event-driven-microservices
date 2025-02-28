package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"gorm.io/gorm"
)

type updateAccountByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewUpdateAccountByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateAccountByID {
	return &updateAccountByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *updateAccountByIDImpl) UpdateAccountByID(ctx context.Context, update entities.UpdateAccount) error {
	if update.NoUpdates() {
		return nil
	}
	updateMap := update.ToUpdateMap()
	if err := handle.db.WithContext(ctx).Model(&entities.Account{}).Where("id = ?", update.ID).Updates(updateMap).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error updating account by id")
		return err
	}
	return nil
}
