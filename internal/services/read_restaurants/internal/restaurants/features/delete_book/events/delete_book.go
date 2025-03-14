package events

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	deletebookcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/commands"
)

type DeleteBook struct {
	*types.MessageImpl
	aggregates.Book
}

func (c *DeleteBook) ToCommand() *deletebookcommands.DeleteBook {
	command := deletebookcommands.DeleteBook(c.Book)
	return &command
}

func NewDeleteBookHandler(
	logger loggers.Logger,
	deleteBookCommand deletebookcommands.DeleteBookHandler,
) *DeleteBookHandle {
	return &DeleteBookHandle{
		logger:            logger,
		deleteBookCommand: deleteBookCommand,
	}
}

type DeleteBookHandle struct {
	logger            loggers.Logger
	deleteBookCommand deletebookcommands.DeleteBookHandler
}

func (h *DeleteBookHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	book, ok := event.Message().(*DeleteBook)
	if !ok {
		h.logger.Error("error in casting DeleteBook event")
		return customizeerrors.BookEventCastingError
	}

	command := book.ToCommand()
	err := h.deleteBookCommand.DeleteBook(ctx, command)
	if err != nil {
		h.logger.Error("error in sending DeleteBook with id: {%s}, error: {%v}", command.Book.ID, err)
		return err
	}
	h.logger.Info("DeleteBook consumer handled.")

	return nil
}
