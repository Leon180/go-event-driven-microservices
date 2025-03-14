package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	createbookcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/commands"
)

type CreateBook struct {
	*types.MessageImpl
	aggregates.Book
}

func (c *CreateBook) ToCommand() *createbookcommands.CreateBook {
	command := createbookcommands.CreateBook(c.Book)
	return &command
}

func NewCreateBookHandler(
	logger loggers.Logger,
	createBookCommand createbookcommands.CreateBookHandler,
) *CreateBookHandle {
	return &CreateBookHandle{
		logger:            logger,
		createBookCommand: createBookCommand,
	}
}

type CreateBookHandle struct {
	logger            loggers.Logger
	createBookCommand createbookcommands.CreateBookHandler
}

func (h *CreateBookHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	book, ok := event.Message().(*CreateBook)
	if !ok {
		h.logger.Error("error in casting CreateBook event")
		return customizeerrors.BookEventCastingError
	}

	command := book.ToCommand()
	err := h.createBookCommand.CreateBook(ctx, command)
	if err != nil {
		h.logger.Error("error in sending CreateBook with id: {%s}, error: {%v}", command.Book.ID, err)
		return err
	}
	h.logger.Info("CreateBook consumer handled.")

	return nil
}
