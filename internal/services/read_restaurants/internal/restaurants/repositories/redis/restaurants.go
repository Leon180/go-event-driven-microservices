package repositoriesredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
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
	var data []byte
	if err := impl.db.Get(ctx, id).Scan(&data); err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to read restaurant full info", err)
		return nil, err
	}
	var restaurant aggregates.Restaurant
	if err := json.Unmarshal(data, &restaurant); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to unmarshal restaurant when reading from redis: ", err)
		return nil, err
	}
	return &restaurant, nil
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
	data, err := json.Marshal(*restaurant)
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to marshal restaurant when setting to redis: ", err)
		return err
	}
	if _, err := impl.db.Set(ctx, restaurant.ID, data, timeOut).Result(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create restaurants: ", err)
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
		if err == redis.Nil {
			return nil
		}
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete restaurant", err)
		return err
	}
	return nil
}
