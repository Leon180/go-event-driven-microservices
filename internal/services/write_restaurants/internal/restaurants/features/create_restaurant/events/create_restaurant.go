package integrationevents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type CreateRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

type CreateRestaurantMessageBuilder interface {
	Build(restaurant *aggregates.Restaurant) *CreateRestaurant
}

func NewCreateRestaurantMessageBuilder(uuidGenerator uuid.UUIDGenerator) CreateRestaurantMessageBuilder {
	return &createRestaurantMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type createRestaurantMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *createRestaurantMessageBuilderImpl) Build(restaurant *aggregates.Restaurant) *CreateRestaurant {
	return &CreateRestaurant{
		Restaurant:  *restaurant,
		MessageImpl: types.NewMessageImpl(b.uuidGenerator.GenerateUUID(), "CreateRestaurant"),
	}
}
