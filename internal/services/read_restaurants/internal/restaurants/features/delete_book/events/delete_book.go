package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	deletebookservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/services"
)

type DeleteBook struct {
	types.Message
	aggregates.Book
}

func NewDeleteBookHandler(
	logger loggers.Logger,
	deleteBookService deletebookservices.DeleteBookHandler,
) *DeleteBookHandle {
	return &DeleteBookHandle{
		logger:            logger,
		deleteBookService: deleteBookService,
	}
}

type DeleteBookHandle struct {
	logger            loggers.Logger
	deleteBookService deletebookservices.DeleteBookHandler
}

func (h *DeleteBookHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*DeleteBook)
	if !ok {
		h.logger.Error("error in casting DeleteBook event")
		return customizeerrors.BookEventCastingError
	}

	err := h.deleteBookService.DeleteBook(ctx, &m.Book)
	if err != nil {
		h.logger.Error("error in sending DeleteBook with id: {%s}, error: {%v}", m.Book.ID, err)
		return err
	}
	h.logger.Info("DeleteBook consumer handled.")

	return nil
}
