package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	updaterestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/commands"
)

type UpdateRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func (c *UpdateRestaurant) ToCommand() *updaterestaurantcommands.UpdateRestaurant {
	command := updaterestaurantcommands.UpdateRestaurant(c.Restaurant)
	return &command
}

func NewUpdateRestaurantHandler(
	logger loggers.Logger,
	updateRestaurantCommand updaterestaurantcommands.UpdateRestaurantHandler,
) *UpdateRestaurantHandle {
	return &UpdateRestaurantHandle{
		logger:                  logger,
		updateRestaurantCommand: updateRestaurantCommand,
	}
}

type UpdateRestaurantHandle struct {
	logger                  loggers.Logger
	updateRestaurantCommand updaterestaurantcommands.UpdateRestaurantHandler
}

func (h *UpdateRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	restaurant, ok := event.Message().(*UpdateRestaurant)
	if !ok {
		h.logger.Error("error in casting UpdateRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	command := restaurant.ToCommand()
	err := h.updateRestaurantCommand.UpdateRestaurant(ctx, command)
	if err != nil {
		h.logger.Error("error in sending UpdateRestaurant with id: {%s}, error: {%v}", command.Restaurant.ID, err)
		return err
	}
	h.logger.Info("UpdateRestaurant consumer handled.")

	return nil
}
