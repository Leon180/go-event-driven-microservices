package repositoriesredis

import (
	"context"
	"encoding/json"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
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
	categories := make(aggregates.Categories, 0)
	l, err := impl.db.HGetAll(ctx, "categories").Result()
	if err != nil {
		return nil, err
	}
	for _, v := range l {
		var category aggregates.Category
		if err := json.Unmarshal([]byte(v), &category); err != nil {
			impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to unmarshal category", err)
			continue
		}
		categories = append(categories, category)
	}
	return categories, nil
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

func (impl *setCategoriesRedisImpl) SetCategories(ctx context.Context, categories aggregates.Categories) error {
	if len(categories) == 0 {
		return nil
	}
	categoryMap := make(map[string]any, len(categories))
	for _, category := range categories {
		jsonBytes, err := json.Marshal(category)
		if err != nil {
			impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to marshal category", err)
			return err
		}
		categoryMap[category.CategoryCode.ToCategory().String()] = string(jsonBytes)
	}
	if err := impl.db.Del(ctx, "categories").Err(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete categories", err)
		return err
	}
	if _, err := impl.db.HSet(ctx, "categories", categoryMap).Result(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to set category", err)
		return err
	}
	return nil
}
