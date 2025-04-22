package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
)

//go:generate mockgen -source=restaurants_mongo.go -destination=./mocks/restaurants_mongo_mock.go -package=mocks

type SearchRestaurantsMongo interface {
	SearchRestaurants(
		ctx context.Context,
		searchRestaurants *dtos.SearchRestaurants,
	) (aggregates.Restaurants, error)
}

type ReadRestaurantsMongo interface {
	ReadRestaurant(ctx context.Context, id string) (*aggregates.Restaurant, error)
}

type ReadRestaurantBranchMongo interface {
	ReadRestaurantBranch(ctx context.Context, branchID string) (*aggregates.Restaurant, error)
}

type UpdateRestaurantsMongo interface {
	CreateRestaurants(ctx context.Context, restaurants aggregates.Restaurants) error
	UpdateRestaurant(ctx context.Context, restaurant *aggregates.Restaurant) error
	DeleteRestaurants(ctx context.Context, ids []string) error
}
