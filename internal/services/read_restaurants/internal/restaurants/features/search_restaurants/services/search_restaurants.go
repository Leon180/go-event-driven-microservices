package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type SearchRestaurants interface {
	SearchRestaurants(
		ctx context.Context,
		req *dtos.SearchRestaurants,
	) (aggregates.Restaurants, error)
}

func NewSearchRestaurants(
	searchRestaurantsFullInfoRepository repositories.SearchRestaurantsFullInfo,
) SearchRestaurants {
	return &searchRestaurantsImpl{searchRestaurantsFullInfoRepository: searchRestaurantsFullInfoRepository}
}

type searchRestaurantsImpl struct {
	searchRestaurantsFullInfoRepository repositories.SearchRestaurantsFullInfo
}

func (handle *searchRestaurantsImpl) SearchRestaurants(
	ctx context.Context,
	req *dtos.SearchRestaurants,
) (aggregates.Restaurants, error) {
	if req == nil {
		return nil, nil
	}
	return handle.searchRestaurantsFullInfoRepository.SearchRestaurantsFullInfo(ctx, req)
}
