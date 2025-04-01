package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	deleterestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
)

type DeleteRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func NewDeleteRestaurantHandler(
	logger loggers.Logger,
	deleteRestaurantService deleterestaurantservices.DeleteRestaurantHandler,
) *DeleteRestaurantHandle {
	return &DeleteRestaurantHandle{
		logger:                  logger,
		deleteRestaurantService: deleteRestaurantService,
	}
}

type DeleteRestaurantHandle struct {
	logger                  loggers.Logger
	deleteRestaurantService deleterestaurantservices.DeleteRestaurantHandler
}

func (h *DeleteRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*DeleteRestaurant)
	if !ok {
		h.logger.Error("error in casting DeleteRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}
	if err := h.deleteRestaurantService.DeleteRestaurant(ctx, &m.Restaurant); err != nil {
		h.logger.Errorf("error in sending DeleteRestaurant with id: {%s}, error: {%v}", m.Restaurant.ID, err)
		return err
	}
	h.logger.Info("DeleteRestaurant consumer handled.")
	return nil
}
