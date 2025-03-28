package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/dtos"
	deleteRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type DeleteRestaurant interface {
	DeleteRestaurant(ctx context.Context, req *featuresdtos.DeleteRestaurantRequest) error
}

func NewDeleteRestaurant(
	updateRestaurantsRepository repositories.UpdateRestaurants,
	readRestaurantsRepository repositories.ReadRestaurants,
	rabbitmqProducer producer.Producer,
	deleteRestaurantMessageBuilder deleteRestaurantEvents.DeleteRestaurantMessageBuilder,
) DeleteRestaurant {
	return &deleteRestaurantImpl{
		updateRestaurantsRepository:    updateRestaurantsRepository,
		readRestaurantsRepository:      readRestaurantsRepository,
		rabbitmqProducer:               rabbitmqProducer,
		deleteRestaurantMessageBuilder: deleteRestaurantMessageBuilder,
	}
}

type deleteRestaurantImpl struct {
	updateRestaurantsRepository    repositories.UpdateRestaurants
	readRestaurantsRepository      repositories.ReadRestaurants
	rabbitmqProducer               producer.Producer
	deleteRestaurantMessageBuilder deleteRestaurantEvents.DeleteRestaurantMessageBuilder
}

func (handle *deleteRestaurantImpl) DeleteRestaurant(
	ctx context.Context,
	req *featuresdtos.DeleteRestaurantRequest,
) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}

	restaurant, err := handle.readRestaurantsRepository.ReadRestaurant(ctx, req.ID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if !restaurant.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	updateRestaurant := entities.UpdateRestaurant{
		ID:           restaurant.ID,
		ActiveStatus: lo.ToPtr(false),
	}
	if err := handle.updateRestaurantsRepository.UpdateRestaurant(ctx, &updateRestaurant); err != nil {
		return err
	}

	// publish delete restaurant event
	re := aggregates.RestaurantEntity(*restaurant)
	message := handle.deleteRestaurantMessageBuilder.Build(re.ToAggregate())
	if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
		return err
	}

	return nil
}
