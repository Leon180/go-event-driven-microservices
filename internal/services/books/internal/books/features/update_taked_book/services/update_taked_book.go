package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/update_taked_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories"
)

type UpdateTakedBook interface {
	UpdateTakedBook(ctx context.Context, req *featuresdtos.UpdateTakedBookRequest) error
}

func NewUpdateTakedBook(
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	readBookRepository repositories.ReadBook,
) UpdateTakedBook {
	return &updateTakedBookImpl{
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		readBookRepository:                   readBookRepository,
	}
}

type updateTakedBookImpl struct {
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	readBookRepository                   repositories.ReadBook
}

func (handle *updateTakedBookImpl) UpdateTakedBook(ctx context.Context, req *featuresdtos.UpdateTakedBookRequest) error {
	if req == nil {
		return nil
	}
	// check if book already exists
	book, err := handle.readBookRepository.ReadBook(ctx, req.BookID)
	if err != nil {
		return err
	}
	if book == nil {
		return customizeerrors.BookNotFoundError
	}

	if req.MobileNumber != book.MobileNumber {
		return customizeerrors.HTTPNoAuthorizationError
	}

	if req.Amount != nil && *req.Amount > book.Capacity {
		return customizeerrors.BookCapacityExceededError
	}

	updateBook := entities.UpdateBook{
		ID:     req.BookID,
		Amount: req.Amount,
		Note:   req.Note,
	}

	// update books
	tx, err := handle.updateBooksWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := tx.UpdateBook(ctx, &updateBook); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
