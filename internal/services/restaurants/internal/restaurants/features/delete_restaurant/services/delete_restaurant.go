package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/delete_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
)

type DeleteRestaurant interface {
	DeleteRestaurant(ctx context.Context, req *featuresdtos.DeleteRestaurantRequest) error
}

func NewDeleteRestaurant(
	restaurantRepository repositories.Restaurants,
) DeleteRestaurant {
	return &deleteRestaurantImpl{
		restaurantRepository: restaurantRepository,
	}
}

type deleteRestaurantImpl struct {
	restaurantRepository repositories.Restaurants
}

func (handle *deleteRestaurantImpl) DeleteRestaurant(ctx context.Context, req *featuresdtos.DeleteRestaurantRequest) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	restaurant, err := handle.restaurantRepository.ReadRestaurant(ctx, req.ID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if !restaurant.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	activeStatus := false
	updateRestaurant := entities.UpdateRestaurant{
		ID:           restaurant.ID,
		ActiveStatus: &activeStatus,
	}
	return handle.restaurantRepository.UpdateRestaurant(ctx, &updateRestaurant)
}
