package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type DeleteRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}
type DeleteRestaurantMessageBuilder interface {
	Build(restaurant *aggregates.Restaurant) *DeleteRestaurant
}

func NewDeleteRestaurantMessageBuilder(uuidGenerator uuid.UUIDGenerator) DeleteRestaurantMessageBuilder {
	return &deleteRestaurantMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type deleteRestaurantMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *deleteRestaurantMessageBuilderImpl) Build(restaurant *aggregates.Restaurant) *DeleteRestaurant {
	return &DeleteRestaurant{
		Restaurant:  *restaurant,
		MessageImpl: types.NewMessageImpl(b.uuidGenerator.GenerateUUID(), "DeleteRestaurant"),
	}
}
