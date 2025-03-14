package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	deleterestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/commands"
)

type DeleteRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func (c *DeleteRestaurant) ToCommand() *deleterestaurantcommands.DeleteRestaurant {
	command := deleterestaurantcommands.DeleteRestaurant(c.Restaurant)
	return &command
}

func NewDeleteRestaurantHandler(
	logger loggers.Logger,
	deleteRestaurantCommand deleterestaurantcommands.DeleteRestaurantHandler,
) *DeleteRestaurantHandle {
	return &DeleteRestaurantHandle{
		logger:                  logger,
		deleteRestaurantCommand: deleteRestaurantCommand,
	}
}

type DeleteRestaurantHandle struct {
	logger                  loggers.Logger
	deleteRestaurantCommand deleterestaurantcommands.DeleteRestaurantHandler
}

func (h *DeleteRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	restaurant, ok := event.Message().(*DeleteRestaurant)
	if !ok {
		h.logger.Error("error in casting DeleteRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	command := restaurant.ToCommand()
	err := h.deleteRestaurantCommand.DeleteRestaurant(ctx, command)
	if err != nil {
		h.logger.Error("error in sending DeleteRestaurant with id: {%s}, error: {%v}", command.Restaurant.ID, err)
		return err
	}
	h.logger.Info("DeleteRestaurant consumer handled.")

	return nil
}
