package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	createrestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/commands"
)

type CreateRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func (c *CreateRestaurant) ToCommand() *createrestaurantcommands.CreateRestaurant {
	command := createrestaurantcommands.CreateRestaurant(c.Restaurant)
	return &command
}

func NewCreateRestaurantHandler(
	logger loggers.Logger,
	createRestaurantCommand createrestaurantcommands.CreateRestaurantHandler,
) *CreateRestaurantHandle {
	return &CreateRestaurantHandle{
		logger:                  logger,
		createRestaurantCommand: createRestaurantCommand,
	}
}

type CreateRestaurantHandle struct {
	logger                  loggers.Logger
	createRestaurantCommand createrestaurantcommands.CreateRestaurantHandler
}

func (h *CreateRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	restaurant, ok := event.Message().(*CreateRestaurant)
	if !ok {
		h.logger.Error("error in casting CreateRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	command := restaurant.ToCommand()
	err := h.createRestaurantCommand.CreateRestaurant(ctx, command)
	if err != nil {
		h.logger.Error("error in sending CreateRestaurant with id: {%s}, error: {%v}", command.Restaurant.ID, err)
		return err
	}
	h.logger.Info("CreateRestaurant consumer handled.")

	return nil
}
