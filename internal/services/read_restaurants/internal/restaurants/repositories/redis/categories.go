package repositoriesredis

import (
	"context"
	"encoding/json"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/redis/go-redis/v9"
)

func NewListCategoriesRedis(
	db *redis.Client,
	contextLogger contextloggers.ContextLogger,
) repositories.ListCategoriesRedis {
	return &listCategoriesRedisImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type listCategoriesRedisImpl struct {
	db            *redis.Client
	contextLogger contextloggers.ContextLogger
}

func (impl *listCategoriesRedisImpl) ListCategories(ctx context.Context) (aggregates.Categories, error) {
	categories := make(documents.Categories, 0)
	l, err := impl.db.HGetAll(ctx, "categories").Result()
	if err != nil {
		return nil, err
	}
	for _, v := range l {
		var category documents.Category
		if err := json.Unmarshal([]byte(v), &category); err != nil {
			impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to unmarshal category", err)
			continue
		}
		categories = append(categories, category)
	}
	return aggregates.CategoryDocuments(categories).ToAggregate(), nil
}

func NewSetCategoriesRedis(
	db *redis.Client,
	contextLogger contextloggers.ContextLogger,
) repositories.SetCategoriesRedis {
	return &setCategoriesRedisImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type setCategoriesRedisImpl struct {
	db            *redis.Client
	contextLogger contextloggers.ContextLogger
}

func (impl *setCategoriesRedisImpl) SetCategory(ctx context.Context, category *aggregates.Category) error {
	if category == nil || category.ID == "" {
		return nil
	}
	categoryDocument := category.ToDocument()
	if _, err := impl.db.HSet(ctx, "categories", *categoryDocument).Result(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to set category", err)
		return err
	}
	return nil
}
