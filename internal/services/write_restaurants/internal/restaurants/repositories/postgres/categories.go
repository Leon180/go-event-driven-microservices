package repostgresespostgres

import (
	"context"

	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"gorm.io/gorm"
)

func NewListCategories(db *gorm.DB, contextLogger contextloggers.ContextLogger) repositories.ListCategories {
	return &listCategoriesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type listCategoriesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *listCategoriesImpl) ListCategories(ctx context.Context) (entities.Categories, error) {
	var categories entities.Categories
	if err := impl.db.WithContext(ctx).Find(&categories).Error; err != nil {
		impl.contextLogger.Error(ctx, "failed to list categories", err)
		return nil, err
	}
	return categories, nil
}
