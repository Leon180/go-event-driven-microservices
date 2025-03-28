package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/dtos"
	restoreRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type RestoreRestaurant interface {
	RestoreRestaurant(ctx context.Context, req *featuresdtos.RestoreRestaurantRequest) error
}

func NewRestoreRestaurant(
	updateRestaurantsRepository repositories.UpdateRestaurants,
	readRestaurantsRepository repositories.ReadRestaurants,
	rabbitmqProducer producer.Producer,
	restoreRestaurantMessageBuilder restoreRestaurantEvents.RestoreRestaurantMessageBuilder,
) RestoreRestaurant {
	return &restoreRestaurantImpl{
		updateRestaurantsRepository:     updateRestaurantsRepository,
		readRestaurantsRepository:       readRestaurantsRepository,
		rabbitmqProducer:                rabbitmqProducer,
		restoreRestaurantMessageBuilder: restoreRestaurantMessageBuilder,
	}
}

type restoreRestaurantImpl struct {
	updateRestaurantsRepository     repositories.UpdateRestaurants
	readRestaurantsRepository       repositories.ReadRestaurants
	rabbitmqProducer                producer.Producer
	restoreRestaurantMessageBuilder restoreRestaurantEvents.RestoreRestaurantMessageBuilder
}

func (handle *restoreRestaurantImpl) RestoreRestaurant(
	ctx context.Context,
	req *featuresdtos.RestoreRestaurantRequest,
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
	if restaurant.IsActive() {
		return customizeerrors.AlreadyActiveError
	}
	updateRestaurant := entities.UpdateRestaurant{
		ID:           restaurant.ID,
		ActiveStatus: lo.ToPtr(true),
	}
	if err := handle.updateRestaurantsRepository.UpdateRestaurant(ctx, &updateRestaurant); err != nil {
		return err
	}

	// publish restore restaurant event
	re := aggregates.RestaurantEntity(*restaurant)
	message := handle.restoreRestaurantMessageBuilder.Build(re.ToAggregate())
	if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
		return err
	}

	return nil
}
