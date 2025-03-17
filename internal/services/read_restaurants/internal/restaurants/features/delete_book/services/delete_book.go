package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type DeleteBookHandler interface {
	DeleteBook(ctx context.Context, command *aggregates.Book) error
}

func NewDeleteBookHandler(
	updateBooksMongo repositories.UpdateBooksMongo,
	readBooksMongo repositories.ReadBooksMongo,
	setBookRedis repositories.SetBookRedis,
) DeleteBookHandler {
	return &deleteBookImpl{
		updateBooksMongo: updateBooksMongo,
		readBooksMongo:   readBooksMongo,
		setBookRedis:     setBookRedis,
	}
}

type deleteBookImpl struct {
	updateBooksMongo repositories.UpdateBooksMongo
	readBooksMongo   repositories.ReadBooksMongo
	setBookRedis     repositories.SetBookRedis
}

func (handle *deleteBookImpl) DeleteBook(ctx context.Context, command *aggregates.Book) error {
	if command == nil {
		return nil
	}
	if command.ID == "" {
		return customizeerrors.InvalidIDError
	}
	book, err := handle.readBooksMongo.ReadBook(ctx, command.ID)
	if err != nil {
		return err
	}
	if book == nil {
		return customizeerrors.BookNotFoundError
	}
	if !book.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}

	aggregate := aggregates.Book(*command)

	err = handle.updateBooksMongo.UpdateBook(ctx, &aggregate)
	if err != nil {
		return err
	}

	err = handle.setBookRedis.DeleteBook(ctx, &aggregate)
	if err != nil {
		return err
	}

	return nil
}
