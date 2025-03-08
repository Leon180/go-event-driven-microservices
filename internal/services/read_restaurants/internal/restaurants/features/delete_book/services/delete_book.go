package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type DeleteBook interface {
	DeleteBook(ctx context.Context, req *featuresdtos.DeleteBookRequest) error
}

func NewDeleteBook(
	updateBooksRepository repositories.UpdateBooks,
	readBooksRepository repositories.ReadBooks,
) DeleteBook {
	return &deleteBookImpl{
		updateBooksRepository: updateBooksRepository,
		readBooksRepository:   readBooksRepository,
	}
}

type deleteBookImpl struct {
	updateBooksRepository repositories.UpdateBooks
	readBooksRepository   repositories.ReadBooks
}

func (handle *deleteBookImpl) DeleteBook(ctx context.Context, req *featuresdtos.DeleteBookRequest) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	book, err := handle.readBooksRepository.ReadBook(ctx, req.ID)
	if err != nil {
		return err
	}
	if book == nil {
		return customizeerrors.BookNotFoundError
	}
	if !book.IsActive() {
		return customizeerrors.AlreadyDeletedError
	}
	updateBook := entities.UpdateBook{
		ID:           book.ID,
		ActiveStatus: lo.ToPtr(false),
	}
	return handle.updateBooksRepository.UpdateBook(ctx, &updateBook)
}
