package repositories

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
)

//go:generate mockgen -source=restaurants_redis.go -destination=./mocks/restaurants_redis_mock.go -package=mocks

type ReadRestaurantsRedis interface {
	ReadRestaurant(ctx context.Context, id string) (*aggregates.Restaurant, error)
}

type SetRestaurantsRedis interface {
	SetRestaurant(ctx context.Context, restaurant *aggregates.Restaurant, timeOut time.Duration) error
	DeleteRestaurant(ctx context.Context, restaurant *aggregates.Restaurant) error
}
