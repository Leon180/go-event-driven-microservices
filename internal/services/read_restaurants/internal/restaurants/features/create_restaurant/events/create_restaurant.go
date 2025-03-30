package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	createrestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/services"
)

type CreateRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func NewCreateRestaurantHandler(
	logger loggers.Logger,
	createRestaurantService createrestaurantservices.CreateRestaurantHandler,
) *CreateRestaurantHandle {
	return &CreateRestaurantHandle{
		logger:                  logger,
		createRestaurantService: createRestaurantService,
	}
}

type CreateRestaurantHandle struct {
	logger                  loggers.Logger
	createRestaurantService createrestaurantservices.CreateRestaurantHandler
}

func (h *CreateRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*CreateRestaurant)
	if !ok {
		h.logger.Error("error in casting CreateRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	err := h.createRestaurantService.CreateRestaurant(ctx, &m.Restaurant)
	if err != nil {
		h.logger.Error("error in sending CreateRestaurant with id: {%s}, error: {%v}", m.Restaurant.ID, err)
		return err
	}
	h.logger.Info("CreateRestaurant consumer handled.")

	return nil
}
