package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
)

type UpdateRestaurant struct {
	types.Message
	aggregates.Restaurant
}

type UpdateRestaurantBuilder interface {
	Build(restaurant aggregates.Restaurant) *UpdateRestaurant
}

func NewUpdateRestaurantBuilder(uuidGenerator uuid.UUIDGenerator) UpdateRestaurantBuilder {
	return &updateRestaurantBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type updateRestaurantBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *updateRestaurantBuilderImpl) Build(restaurant aggregates.Restaurant) *UpdateRestaurant {
	return &UpdateRestaurant{
		Restaurant: restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "UpdateRestaurant"),
	}
}
