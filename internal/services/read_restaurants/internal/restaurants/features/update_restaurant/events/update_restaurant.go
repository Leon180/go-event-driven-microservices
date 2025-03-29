package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	updaterestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/services"
)

type UpdateRestaurant struct {
	types.Message
	aggregates.Restaurant
}

func NewUpdateRestaurantHandler(
	logger loggers.Logger,
	updateRestaurantService updaterestaurantservices.UpdateRestaurantHandler,
) *UpdateRestaurantHandle {
	return &UpdateRestaurantHandle{
		logger:                  logger,
		updateRestaurantService: updateRestaurantService,
	}
}

type UpdateRestaurantHandle struct {
	logger                  loggers.Logger
	updateRestaurantService updaterestaurantservices.UpdateRestaurantHandler
}

func (h *UpdateRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*UpdateRestaurant)
	if !ok {
		h.logger.Error("error in casting UpdateRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	err := h.updateRestaurantService.UpdateRestaurant(ctx, &m.Restaurant)
	if err != nil {
		h.logger.Error("error in sending UpdateRestaurant with id: {%s}, error: {%v}", m.Restaurant.ID, err)
		return err
	}
	h.logger.Info("UpdateRestaurant consumer handled.")

	return nil
}
