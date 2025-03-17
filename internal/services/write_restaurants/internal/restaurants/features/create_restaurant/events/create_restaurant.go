package integrationevents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
)

type CreateRestaurant struct {
	types.Message
	aggregates.Restaurant
}

type CreateRestaurantBuilder interface {
	Build(restaurant aggregates.Restaurant) *CreateRestaurant
}

func NewCreateRestaurantBuilder(uuidGenerator uuid.UUIDGenerator) CreateRestaurantBuilder {
	return &createRestaurantBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type createRestaurantBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *createRestaurantBuilderImpl) Build(restaurant aggregates.Restaurant) *CreateRestaurant {
	return &CreateRestaurant{
		Restaurant: restaurant,
		Message:    types.NewMessage(b.uuidGenerator.GenerateUUID(), "CreateRestaurant"),
	}
}
