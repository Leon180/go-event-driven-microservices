package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	synccategorieservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/sync_categories/services"
)

type SyncCategories struct {
	*types.MessageImpl
	aggregates.Categories
}

func NewSyncCategoriesHandler(
	logger loggers.Logger,
	syncCategoriesService synccategorieservices.SyncCategoriesHandler,
) *SyncCategoriesHandle {
	return &SyncCategoriesHandle{
		logger:                logger,
		syncCategoriesService: syncCategoriesService,
	}
}

type SyncCategoriesHandle struct {
	logger                loggers.Logger
	syncCategoriesService synccategorieservices.SyncCategoriesHandler
}

func (h *SyncCategoriesHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*SyncCategories)
	if !ok {
		h.logger.Error("error in casting SyncCategories event")
		return customizeerrors.BookEventCastingError
	}
	if err := h.syncCategoriesService.SyncCategories(ctx, m.Categories); err != nil {
		h.logger.Errorf("error in sending SyncCategories, error: {%v}", err)
		return err
	}
	h.logger.Info("SyncCategories consumer handled.")
	return nil
}
