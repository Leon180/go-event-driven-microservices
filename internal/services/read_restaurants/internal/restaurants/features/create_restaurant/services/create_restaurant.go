package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type CreateRestaurantHandler interface {
	CreateRestaurant(ctx context.Context, aggregate *aggregates.Restaurant) error
}

func NewCreateRestaurantHandler(
	redisConfig *redisdb.RedisConfig,
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo,
	readRestaurantsMongo repositories.ReadRestaurantsMongo,
	setRestaurantsRedis repositories.SetRestaurantsRedis,
) CreateRestaurantHandler {
	return &createRestaurantImpl{
		redisConfig:            redisConfig,
		updateRestaurantsMongo: updateRestaurantsMongo,
		readRestaurantsMongo:   readRestaurantsMongo,
		setRestaurantsRedis:    setRestaurantsRedis,
	}
}

type createRestaurantImpl struct {
	redisConfig            *redisdb.RedisConfig
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo
	readRestaurantsMongo   repositories.ReadRestaurantsMongo
	setRestaurantsRedis    repositories.SetRestaurantsRedis
}

func (handle *createRestaurantImpl) CreateRestaurant(ctx context.Context, aggregate *aggregates.Restaurant) error {
	if aggregate == nil {
		return nil
	}
	// check if restaurant already exists
	existed, err := handle.readRestaurantsMongo.ReadRestaurant(ctx, aggregate.ID)
	if err != nil {
		return err
	}
	if existed != nil {
		if existed.ActiveStatus {
			return customizeerrors.RestaurantAlreadyExistsError
		}
		return customizeerrors.RestaurantAlreadyExistsButInactiveError
	}
	if err = handle.updateRestaurantsMongo.CreateRestaurants(ctx, []aggregates.Restaurant{*aggregate}); err != nil {
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
