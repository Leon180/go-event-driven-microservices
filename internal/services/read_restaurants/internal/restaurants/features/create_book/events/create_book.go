package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	createbookservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/services"
)

type CreateBook struct {
	*types.MessageImpl
	aggregates.Book
}

func NewCreateBookHandler(
	logger loggers.Logger,
	createBookService createbookservices.CreateBookHandler,
) *CreateBookHandle {
	return &CreateBookHandle{
		logger:            logger,
		createBookService: createBookService,
	}
}

type CreateBookHandle struct {
	logger            loggers.Logger
	createBookService createbookservices.CreateBookHandler
}

func (h *CreateBookHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	m, ok := event.Message().(*CreateBook)
	if !ok {
		h.logger.Error("error in casting CreateBook event")
		return customizeerrors.BookEventCastingError
	}

	err := h.createBookService.CreateBook(ctx, &m.Book)
	if err != nil {
		h.logger.Error("error in sending CreateBook with id: {%s}, error: {%v}", m.Book.ID, err)
		return err
	}
	h.logger.Info("CreateBook consumer handled.")

	return nil
}
