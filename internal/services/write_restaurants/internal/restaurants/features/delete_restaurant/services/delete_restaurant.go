package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
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

	existed, err := handle.readRestaurantsRepository.ReadRestaurantFullInfo(ctx, req.ID)
	if err != nil {
		return err
	}
	if existed == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if !existed.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	updateRestaurant := entities.UpdateRestaurant{
		ID:           existed.ID,
		ActiveStatus: lo.ToPtr(false),
	}
	if err := handle.updateRestaurantsRepository.UpdateRestaurant(ctx, &updateRestaurant); err != nil {
		return err
	}

	// publish delete restaurant event
	message := handle.deleteRestaurantMessageBuilder.Build(existed)
	if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
		return err
	}

	return nil
}
