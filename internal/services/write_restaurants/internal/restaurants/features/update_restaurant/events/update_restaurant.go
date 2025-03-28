package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type UpdateRestaurant struct {
	types.Message
	aggregates.Restaurant
}

type UpdateRestaurantMessageBuilder interface {
	Build(restaurant *aggregates.Restaurant) *UpdateRestaurant
}

func NewUpdateRestaurantMessageBuilder(uuidGenerator uuid.UUIDGenerator) UpdateRestaurantMessageBuilder {
	return &updateRestaurantMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type updateRestaurantMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *updateRestaurantMessageBuilderImpl) Build(restaurant *aggregates.Restaurant) *UpdateRestaurant {
	return &UpdateRestaurant{
		Restaurant: *restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "UpdateRestaurant"),
	}
}
