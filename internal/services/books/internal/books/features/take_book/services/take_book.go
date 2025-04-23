package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/take_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories"
	"github.com/samber/lo"
)

type TakeBook interface {
	TakeBook(ctx context.Context, req *featuresdtos.TakeBookRequest) error
}

func NewTakeBook(
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	readBookRepository repositories.ReadBook,
) TakeBook {
	return &takeBookImpl{
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		readBookRepository:                   readBookRepository,
	}
}

type takeBookImpl struct {
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	readBookRepository                   repositories.ReadBook
}

func (handle *takeBookImpl) TakeBook(ctx context.Context, req *featuresdtos.TakeBookRequest) error {
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

	if req.Amount > book.Capacity {
		return customizeerrors.BookCapacityExceededError
	}

	updateBook := entities.UpdateBook{
		ID:           req.BookID,
		Booked:       lo.ToPtr(true),
		Amount:       lo.ToPtr(req.Amount),
		MobileNumber: lo.ToPtr(req.MobileNumber),
		Note:         lo.ToPtr(req.Note),
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
