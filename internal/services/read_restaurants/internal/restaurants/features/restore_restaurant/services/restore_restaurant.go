package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type RestoreRestaurantHandler interface {
	RestoreRestaurant(ctx context.Context, aggregate *aggregates.Restaurant) error
}

func NewRestoreRestaurantHandler(
	redisConfig *redisdb.RedisConfig,
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo,
	readRestaurantsMongo repositories.ReadRestaurantsMongo,
	setRestaurantsRedis repositories.SetRestaurantsRedis,
) RestoreRestaurantHandler {
	return &restoreRestaurantImpl{
		redisConfig:            redisConfig,
		updateRestaurantsMongo: updateRestaurantsMongo,
		readRestaurantsMongo:   readRestaurantsMongo,
		setRestaurantsRedis:    setRestaurantsRedis,
	}
}

type restoreRestaurantImpl struct {
	redisConfig            *redisdb.RedisConfig
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo
	readRestaurantsMongo   repositories.ReadRestaurantsMongo
	setRestaurantsRedis    repositories.SetRestaurantsRedis
}

func (handle *restoreRestaurantImpl) RestoreRestaurant(
	ctx context.Context,
	aggregate *aggregates.Restaurant,
) error {
	if aggregate == nil {
		return nil
	}
	if aggregate.ID == "" {
		return customizeerrors.InvalidIDError
	}
	existed, err := handle.readRestaurantsMongo.ReadRestaurant(ctx, aggregate.ID)
	if err != nil {
		return err
	}
	if existed == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if existed.IsActive() {
		return customizeerrors.AlreadyActiveError
	}
	aggregate.ActiveStatus = true
	if err = handle.updateRestaurantsMongo.UpdateRestaurant(ctx, aggregate); err != nil {
		return err
	}
	if err = handle.setRestaurantsRedis.SetRestaurant(ctx, aggregate, time.Duration(handle.redisConfig.CacheTimeOut)*time.Second); err != nil {
		return err
	}
	return nil
}
