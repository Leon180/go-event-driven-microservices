package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
)

type DeleteRestaurant struct {
	types.Message
	aggregates.Restaurant
}

type DeleteRestaurantBuilder interface {
	Build(restaurant aggregates.Restaurant) *DeleteRestaurant
}

func NewDeleteRestaurantBuilder(uuidGenerator uuid.UUIDGenerator) DeleteRestaurantBuilder {
	return &deleteRestaurantBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type deleteRestaurantBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *deleteRestaurantBuilderImpl) Build(restaurant aggregates.Restaurant) *DeleteRestaurant {
	return &DeleteRestaurant{
		Restaurant: restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "DeleteRestaurant"),
	}
}
