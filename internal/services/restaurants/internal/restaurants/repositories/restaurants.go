package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
)

//go:generate mockgen -source=restaurants.go -destination=./mocks/restaurants_mock.go -package=mocks

type CreateRestaurants interface {
	CreateRestaurants(ctx context.Context, restaurant entities.Restaurants) error
}

type ReadRestaurants interface {
	ReadRestaurants(ctx context.Context, ids ...string) (entities.Restaurants, error)
}

type UpdateRestaurant interface {
	UpdateRestaurant(ctx context.Context, updateRestaurant entities.UpdateRestaurant) error
}

type DeleteRestaurant interface {
	DeleteRestaurant(ctx context.Context, id string) error
}
