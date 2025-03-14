package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	restorerestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/commands"
)

type RestoreRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func (c *RestoreRestaurant) ToCommand() *restorerestaurantcommands.RestoreRestaurant {
	command := restorerestaurantcommands.RestoreRestaurant(c.Restaurant)
	return &command
}

func NewRestoreRestaurantHandler(
	logger loggers.Logger,
	restoreRestaurantCommand restorerestaurantcommands.RestoreRestaurantHandler,
) *RestoreRestaurantHandle {
	return &RestoreRestaurantHandle{
		logger:                   logger,
		restoreRestaurantCommand: restoreRestaurantCommand,
	}
}

type RestoreRestaurantHandle struct {
	logger                   loggers.Logger
	restoreRestaurantCommand restorerestaurantcommands.RestoreRestaurantHandler
}

func (h *RestoreRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	restaurant, ok := event.Message().(*RestoreRestaurant)
	if !ok {
		h.logger.Error("error in casting RestoreRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	command := restaurant.ToCommand()
	err := h.restoreRestaurantCommand.RestoreRestaurant(ctx, command)
	if err != nil {
		h.logger.Error("error in sending RestoreRestaurant with id: {%s}, error: {%v}", command.Restaurant.ID, err)
		return err
	}
	h.logger.Info("RestoreRestaurant consumer handled.")

	return nil
}
