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

func (handle *deleteBookImpl) DeleteBook(ctx context.Context, aggregate *aggregates.Book) error {
	if aggregate == nil {
		return nil
	}
	if aggregate.ID == "" {
		return customizeerrors.InvalidIDError
	}
	existed, err := handle.readBooksMongo.ReadBook(ctx, aggregate.ID)
	if err != nil {
		return err
	}
	if existed == nil {
		return customizeerrors.BookNotFoundError
	}
	if !existed.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	aggregate.ActiveStatus = false
	if err = handle.updateBooksMongo.UpdateBook(ctx, aggregate); err != nil {
		return err
	}
	if err = handle.setBookRedis.DeleteBook(ctx, aggregate); err != nil {
		return err
	}
	return nil
}
