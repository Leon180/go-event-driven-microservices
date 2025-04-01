package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
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

	existed, err := handle.readRestaurantsRepository.ReadRestaurantFullInfo(ctx, req.ID)
	if err != nil {
		return err
	}
	if existed == nil {
		return customizeerrors.RestaurantNotFoundError
	}
	if existed.IsActive() {
		return customizeerrors.AlreadyActiveError
	}
	updateRestaurant := entities.UpdateRestaurant{
		ID:           existed.ID,
		ActiveStatus: lo.ToPtr(true),
	}
	if err := handle.updateRestaurantsRepository.UpdateRestaurant(ctx, &updateRestaurant); err != nil {
		return err
	}

	// publish restore restaurant event
	message := handle.restoreRestaurantMessageBuilder.Build(existed)
	if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
		return err
	}

	return nil
}
