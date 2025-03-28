package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type RestoreRestaurant struct {
	types.Message
	aggregates.Restaurant
}

type RestoreRestaurantMessageBuilder interface {
	Build(restaurant *aggregates.Restaurant) *RestoreRestaurant
}

func NewRestoreRestaurantMessageBuilder(uuidGenerator uuid.UUIDGenerator) RestoreRestaurantMessageBuilder {
	return &restoreRestaurantMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type restoreRestaurantMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *restoreRestaurantMessageBuilderImpl) Build(restaurant *aggregates.Restaurant) *RestoreRestaurant {
	return &RestoreRestaurant{
		Restaurant: *restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "RestoreRestaurant"),
	}
}
