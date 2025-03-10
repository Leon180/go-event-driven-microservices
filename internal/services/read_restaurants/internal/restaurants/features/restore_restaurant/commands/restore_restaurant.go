package commands

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type RestoreRestaurant aggregates.Restaurant

type RestoreRestaurantHandler interface {
	RestoreRestaurant(ctx context.Context, command *RestoreRestaurant) error
}

func NewRestoreRestaurantHandler(
	redisConfig redisdb.RedisConfig,
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
	redisConfig            redisdb.RedisConfig
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo
	readRestaurantsMongo   repositories.ReadRestaurantsMongo
	setRestaurantsRedis    repositories.SetRestaurantsRedis
}

func (handle *restoreRestaurantImpl) RestoreRestaurant(
	ctx context.Context,
	command *RestoreRestaurant,
) error {
	if command == nil {
		return nil
	}
	if command.ID == "" {
		return customizeerrors.InvalidIDError
	}
	restaurant, err := handle.readRestaurantsMongo.ReadRestaurant(ctx, command.ID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if restaurant.IsActive() {
		return customizeerrors.AlreadyActiveError
	}
	aggregate := aggregates.Restaurant(*command)
	if err := handle.updateRestaurantsMongo.UpdateRestaurant(ctx, &aggregate); err != nil {
		return err
	}
	if err := handle.setRestaurantsRedis.SetRestaurant(ctx, &aggregate, time.Duration(handle.redisConfig.CacheTimeOut)*time.Second); err != nil {
		return err
	}
	return nil
}
