package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/restore_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type RestoreRestaurant interface {
	RestoreRestaurant(ctx context.Context, req *featuresdtos.RestoreRestaurantRequest) error
}

func NewRestoreRestaurant(
	updateRestaurantsRepository repositories.UpdateRestaurants,
	readRestaurantsRepository repositories.ReadRestaurants,
) RestoreRestaurant {
	return &restoreRestaurantImpl{
		updateRestaurantsRepository: updateRestaurantsRepository,
		readRestaurantsRepository:   readRestaurantsRepository,
	}
}

type restoreRestaurantImpl struct {
	updateRestaurantsRepository repositories.UpdateRestaurants
	readRestaurantsRepository   repositories.ReadRestaurants
}

func (handle *restoreRestaurantImpl) RestoreRestaurant(
	ctx context.Context,
	req *featuresdtos.RestoreRestaurantRequest,
) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	restaurant, err := handle.readRestaurantsRepository.ReadRestaurant(ctx, req.ID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if restaurant.IsActive() {
		return customizeerrors.AlreadyActiveError
	}
	updateRestaurant := entities.UpdateRestaurant{
		ID:           restaurant.ID,
		ActiveStatus: lo.ToPtr(true),
	}
	return handle.updateRestaurantsRepository.UpdateRestaurant(ctx, &updateRestaurant)
}
