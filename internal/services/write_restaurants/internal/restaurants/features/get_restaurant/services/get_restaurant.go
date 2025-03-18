package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/get_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
)

type GetRestaurant interface {
	GetRestaurant(
		ctx context.Context,
		req *featuresdtos.GetRestaurantRequest,
	) (*aggregates.Restaurant, error)
}

func NewGetRestaurant(
	readRestaurantsRepository repositories.ReadRestaurants,
) GetRestaurant {
	return &getRestaurantImpl{readRestaurantsRepository: readRestaurantsRepository}
}

type getRestaurantImpl struct {
	readRestaurantsRepository repositories.ReadRestaurants
}

func (handle *getRestaurantImpl) GetRestaurant(
	ctx context.Context,
	req *featuresdtos.GetRestaurantRequest,
) (*aggregates.Restaurant, error) {
	if req == nil {
		return nil, nil
	}
	if req.ID == "" {
		return nil, customizeerrors.InvalidIDError
	}
	restaurant, err := handle.readRestaurantsRepository.ReadRestaurantFullInfo(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if restaurant == nil {
		return nil, customizeerrors.RestaurantNotFoundError
	}
	return restaurant, nil
}
