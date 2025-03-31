package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type DeleteRestaurantHandler interface {
	DeleteRestaurant(ctx context.Context, aggregate *aggregates.Restaurant) error
}

func NewDeleteRestaurantHandler(
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
	aggregate *aggregates.Restaurant,
) error {
	if aggregate == nil {
		return nil
	}
	if aggregate.ID == "" {
		return customizeerrors.InvalidIDError
	}
	restaurant, err := handle.readRestaurantsMongo.ReadRestaurant(ctx, aggregate.ID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if !restaurant.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	if err = handle.updateRestaurantsMongo.UpdateRestaurant(ctx, aggregate); err != nil {
		return err
	}
	if err = handle.setRestaurantsRedis.DeleteRestaurant(ctx, aggregate); err != nil {
		return err
	}
	return nil
}
