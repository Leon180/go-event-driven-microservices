package commands

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type DeleteRestaurant aggregates.Restaurant

type DeleteRestaurantHandler interface {
	DeleteRestaurant(ctx context.Context, command *DeleteRestaurant) error
}

func NewDeleteRestaurant(
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo,
	readRestaurantsMongo repositories.ReadRestaurantsMongo,
	setRestaurantsRedis repositories.SetRestaurantsRedis,
) DeleteRestaurantHandler {
	return &deleteRestaurantImpl{
		updateRestaurantsMongo: updateRestaurantsMongo,
		readRestaurantsMongo:   readRestaurantsMongo,
		setRestaurantsRedis:    setRestaurantsRedis,
	}
}

type deleteRestaurantImpl struct {
	updateRestaurantsMongo repositories.UpdateRestaurantsMongo
	readRestaurantsMongo   repositories.ReadRestaurantsMongo
	setRestaurantsRedis    repositories.SetRestaurantsRedis
}

func (handle *deleteRestaurantImpl) DeleteRestaurant(
	ctx context.Context,
	command *DeleteRestaurant,
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
	if !restaurant.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	aggregate := aggregates.Restaurant(*command)
	if err := handle.updateRestaurantsMongo.UpdateRestaurant(ctx, &aggregate); err != nil {
		return err
	}
	if err := handle.setRestaurantsRedis.DeleteRestaurant(ctx, &aggregate); err != nil {
		return err
	}
	return nil
}
