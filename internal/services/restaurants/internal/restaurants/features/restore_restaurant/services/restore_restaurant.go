package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/restore_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
)

type RestoreRestaurant interface {
	RestoreRestaurant(ctx context.Context, req *featuresdtos.RestoreRestaurantRequest) error
}

func NewRestoreRestaurant(
	restaurantRepository repositories.Restaurants,
) RestoreRestaurant {
	return &restoreRestaurantImpl{
		restaurantRepository: restaurantRepository,
	}
}

type restoreRestaurantImpl struct {
	restaurantRepository repositories.Restaurants
}

func (handle *restoreRestaurantImpl) RestoreRestaurant(ctx context.Context, req *featuresdtos.RestoreRestaurantRequest) error {
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
	if restaurant.IsActive() {
		return customizeerrors.AlreadyActiveError
	}
	activeStatus := true
	updateRestaurant := entities.UpdateRestaurant{
		ID:           restaurant.ID,
		ActiveStatus: &activeStatus,
	}
	return handle.restaurantRepository.UpdateRestaurant(ctx, &updateRestaurant)
}
