package integrationevents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type CreateRestaurant struct {
	types.Message
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
		Restaurant: *restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "CreateRestaurant"),
	}
}
