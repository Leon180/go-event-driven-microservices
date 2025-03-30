package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	restorerestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/services"
)

type RestoreRestaurant struct {
	*types.MessageImpl
	aggregates.Restaurant
}

func NewRestoreRestaurantHandler(
	logger loggers.Logger,
	restoreRestaurantService restorerestaurantservices.RestoreRestaurantHandler,
) *RestoreRestaurantHandle {
	return &RestoreRestaurantHandle{
		logger:                   logger,
		restoreRestaurantService: restoreRestaurantService,
	}
}

type RestoreRestaurantHandle struct {
	logger                   loggers.Logger
	restoreRestaurantService restorerestaurantservices.RestoreRestaurantHandler
}

func (h *RestoreRestaurantHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*RestoreRestaurant)
	if !ok {
		h.logger.Error("error in casting RestoreRestaurant event")
		return customizeerrors.RestaurantEventCastingError
	}

	err := h.restoreRestaurantService.RestoreRestaurant(ctx, &m.Restaurant)
	if err != nil {
		h.logger.Error("error in sending RestoreRestaurant with id: {%s}, error: {%v}", m.Restaurant.ID, err)
		return err
	}
	h.logger.Info("RestoreRestaurant consumer handled.")

	return nil
}
