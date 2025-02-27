package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/repositories"
	"gorm.io/gorm"
)

type deleteAccountByIDImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewDeleteAccountByID(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.DeleteAccountByID {
	return &deleteAccountByIDImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *deleteAccountByIDImpl) DeleteAccountByID(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Delete(&entities.Account{}).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error deleting account by id")
		return err
	}
	return nil
}
