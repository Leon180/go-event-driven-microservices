package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
)

type RestoreRestaurant struct {
	types.Message
	aggregates.Restaurant
}

type RestoreRestaurantBuilder interface {
	Build(restaurant aggregates.Restaurant) *RestoreRestaurant
}

func NewRestoreRestaurantBuilder(uuidGenerator uuid.UUIDGenerator) RestoreRestaurantBuilder {
	return &restoreRestaurantBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type restoreRestaurantBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *restoreRestaurantBuilderImpl) Build(restaurant aggregates.Restaurant) *RestoreRestaurant {
	return &RestoreRestaurant{
		Restaurant: restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "RestoreRestaurant"),
	}
}
