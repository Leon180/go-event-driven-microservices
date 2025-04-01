package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type UpdateRestaurantHandler interface {
	UpdateRestaurant(ctx context.Context, aggregate *aggregates.Restaurant) error
}

func NewUpdateRestaurantHandler(
	redisConfig *redisdb.RedisConfig,
	readRestaurantsRepository repositories.ReadRestaurantsMongo,
	updateRestaurantsRepository repositories.UpdateRestaurantsMongo,
	readRestaurantsRedis repositories.ReadRestaurantsRedis,
	setRestaurantsRedis repositories.SetRestaurantsRedis,
) UpdateRestaurantHandler {
	return &updateRestaurantImpl{
		redisConfig:                 redisConfig,
		readRestaurantsRepository:   readRestaurantsRepository,
		updateRestaurantsRepository: updateRestaurantsRepository,
		readRestaurantsRedis:        readRestaurantsRedis,
		setRestaurantsRedis:         setRestaurantsRedis,
	}
}

type updateRestaurantImpl struct {
	redisConfig                 *redisdb.RedisConfig
	readRestaurantsRepository   repositories.ReadRestaurantsMongo
	updateRestaurantsRepository repositories.UpdateRestaurantsMongo
	readRestaurantsRedis        repositories.ReadRestaurantsRedis
	setRestaurantsRedis         repositories.SetRestaurantsRedis
}

func (handle *updateRestaurantImpl) UpdateRestaurant(ctx context.Context, aggregate *aggregates.Restaurant) error {
	if aggregate == nil {
		return nil
	}
	// check if restaurant exists
	existed, err := handle.readRestaurantsRepository.ReadRestaurant(ctx, aggregate.ID)
	if err != nil {
		return err
	}
	if existed == nil {
		return customizeerrors.RestaurantNotFoundError
	}

	if err = handle.updateRestaurantsRepository.UpdateRestaurant(ctx, aggregate); err != nil {
		return err
	}

	if err = handle.setRestaurantsRedis.SetRestaurant(
		ctx,
		aggregate,
		time.Duration(handle.redisConfig.CacheTimeOut)*time.Second,
	); err != nil {
		return err
	}

	return nil
}
