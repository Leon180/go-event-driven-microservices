package repositoriesredis

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/redis/go-redis/v9"
)

func NewReadRestaurantsRedis(
	db *redis.Client,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadRestaurantsRedis {
	return &ReadRestaurantsRedisImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type ReadRestaurantsRedisImpl struct {
	db            *redis.Client
	contextLogger contextloggers.ContextLogger
}

func (impl *ReadRestaurantsRedisImpl) ReadRestaurant(
	ctx context.Context,
	id string,
) (*aggregates.Restaurant, error) {
	if id == "" {
		return nil, nil
	}
	var restaurant documents.Restaurant
	if err := impl.db.Get(ctx, id).Scan(&restaurant); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to read restaurant full info", err)
		return nil, err
	}
	aggregate := aggregates.RestaurantDocument(restaurant)
	return aggregate.ToAggregate(), nil
}

// Set Restaurants Impl
func NewSetRestaurantsRedis(
	db *redis.Client,
	contextLogger contextloggers.ContextLogger,
) repositories.SetRestaurantsRedis {
	return &setRestaurantsRedisImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type setRestaurantsRedisImpl struct {
	db            *redis.Client
	contextLogger contextloggers.ContextLogger
}

func (impl *setRestaurantsRedisImpl) SetRestaurant(
	ctx context.Context,
	restaurant *aggregates.Restaurant,
	timeOut time.Duration,
) error {
	if restaurant == nil || restaurant.ID == "" {
		return nil
	}
	if _, err := impl.db.Set(ctx, restaurant.ID, restaurant, timeOut).Result(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create restaurants", err)
		return err
	}
	return nil
}

func (impl *setRestaurantsRedisImpl) DeleteRestaurant(
	ctx context.Context,
	restaurant *aggregates.Restaurant,
) error {
	if restaurant == nil || restaurant.ID == "" {
		return nil
	}
	if err := impl.db.Del(ctx, restaurant.ID).Err(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete restaurant", err)
		return err
	}
	return nil
}
