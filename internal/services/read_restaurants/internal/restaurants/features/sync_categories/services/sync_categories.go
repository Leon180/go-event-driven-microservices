package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type SyncCategoriesHandler interface {
	SyncCategories(ctx context.Context, aggregate aggregates.Categories) error
}

func NewSyncCategoriesHandler(
	redisConfig redisdb.RedisConfig,
	syncCategoriesMongo repositories.SyncCategoriesMongo,
	setCategoriesRedis repositories.SetCategoriesRedis,
) SyncCategoriesHandler {
	return &syncCategoriesImpl{
		redisConfig:         redisConfig,
		syncCategoriesMongo: syncCategoriesMongo,
		setCategoriesRedis:  setCategoriesRedis,
	}
}

type syncCategoriesImpl struct {
	redisConfig         redisdb.RedisConfig
	syncCategoriesMongo repositories.SyncCategoriesMongo
	setCategoriesRedis  repositories.SetCategoriesRedis
}

func (handle *syncCategoriesImpl) SyncCategories(ctx context.Context, aggregate aggregates.Categories) error {
	if aggregate == nil {
		return nil
	}

	if err := handle.syncCategoriesMongo.SyncCategories(ctx, aggregate); err != nil {
		return err
	}

	if err := handle.setCategoriesRedis.SetCategories(ctx, aggregate); err != nil {
		return err
	}

	return nil
}
