package commands

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type GetRestaurant struct {
	ID string
}

type GetRestaurantHandler interface {
	GetRestaurant(
		ctx context.Context,
		command *GetRestaurant,
	) (*aggregates.Restaurant, error)
}

func NewGetRestaurantHandler(
	redisConfig redisdb.RedisConfig,
	readRestaurantsMongo repositories.ReadRestaurantsMongo,
	readRestaurantsRedis repositories.ReadRestaurantsRedis,
	setRestaurantsRedis repositories.SetRestaurantsRedis,
) GetRestaurantHandler {
	return &getRestaurantImpl{
		redisConfig:          redisConfig,
		readRestaurantsMongo: readRestaurantsMongo,
		readRestaurantsRedis: readRestaurantsRedis,
		setRestaurantsRedis:  setRestaurantsRedis,
	}
}

type getRestaurantImpl struct {
	redisConfig          redisdb.RedisConfig
	readRestaurantsMongo repositories.ReadRestaurantsMongo
	readRestaurantsRedis repositories.ReadRestaurantsRedis
	setRestaurantsRedis  repositories.SetRestaurantsRedis
}

func (handle *getRestaurantImpl) GetRestaurant(
	ctx context.Context,
	command *GetRestaurant,
) (*aggregates.Restaurant, error) {
	if command == nil {
		return nil, nil
	}
	if command.ID == "" {
		return nil, customizeerrors.InvalidIDError
	}
	restaurant, err := handle.readRestaurantsRedis.ReadRestaurant(ctx, command.ID)
	if err != nil {
		return nil, err
	}
	if restaurant != nil {
		return restaurant, nil
	}
	restaurant, err = handle.readRestaurantsMongo.ReadRestaurant(ctx, command.ID)
	if err != nil {
		return nil, err
	}
	if restaurant == nil {
		return nil, customizeerrors.RestaurantNotFoundError
	}
	err = handle.setRestaurantsRedis.SetRestaurant(
		ctx,
		restaurant,
		time.Duration(handle.redisConfig.CacheTimeOut)*time.Second,
	)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
