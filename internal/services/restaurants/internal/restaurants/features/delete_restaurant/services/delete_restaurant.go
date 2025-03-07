package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/delete_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type DeleteRestaurant interface {
	DeleteRestaurant(ctx context.Context, req *featuresdtos.DeleteRestaurantRequest) error
}

func NewDeleteRestaurant(
	updateRestaurantsRepository repositories.UpdateRestaurants,
	readRestaurantsRepository repositories.ReadRestaurants,
) DeleteRestaurant {
	return &deleteRestaurantImpl{
		updateRestaurantsRepository: updateRestaurantsRepository,
		readRestaurantsRepository:   readRestaurantsRepository,
	}
}

type deleteRestaurantImpl struct {
	updateRestaurantsRepository repositories.UpdateRestaurants
	readRestaurantsRepository   repositories.ReadRestaurants
}

func (handle *deleteRestaurantImpl) DeleteRestaurant(ctx context.Context, req *featuresdtos.DeleteRestaurantRequest) error {
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
	if !restaurant.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	updateRestaurant := entities.UpdateRestaurant{
		ID:           restaurant.ID,
		ActiveStatus: lo.ToPtr(false),
	}
	return handle.updateRestaurantsRepository.UpdateRestaurant(ctx, &updateRestaurant)
}
