package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type CreateBook interface {
	CreateBook(ctx context.Context, req *dtos.Book) error
}

func NewCreateBook(
	uuidGenerator uuid.UUIDGenerator,
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo,
) CreateBook {
	return &createBookImpl{
		uuidGenerator:                        uuidGenerator,
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		searchBooksFullInfoRepository:        searchBooksFullInfoRepository,
	}
}

type createBookImpl struct {
	uuidGenerator                        uuid.UUIDGenerator
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	searchBooksFullInfoRepository        repositories.SearchBooksFullInfo
}

func (handle *createBookImpl) CreateBook(ctx context.Context, req *dtos.Book) error {
	if req == nil {
		return nil
	}

	// check if book already exists
	books, err := handle.searchBooksFullInfoRepository.SearchBooksFullInfo(ctx, &dtos.SearchBooks{
		TableID:     &req.TableID,
		AvailableID: &req.AvailableID,
	})
	if err != nil {
		return err
	}
	if lo.ContainsBy(books, func(book aggregates.Book) bool {
		return book.ActiveStatus
	}) {
		return customizeerrors.BookAlreadyExistsError
	}

	if lo.ContainsBy(books, func(book aggregates.Book) bool {
		return !book.ActiveStatus
	}) {
		return customizeerrors.BookAlreadyExistsButInactiveError
	}

	// build book create entities by aggregate
	bookDTOAggregateBuilder := aggregates.NewBookDTOAggregateBuilder(handle.uuidGenerator)
	err = bookDTOAggregateBuilder.SaveBook(req)
	if err != nil {
		return err
	}
	editEntities := bookDTOAggregateBuilder.GetEditEntities()
	if len(editEntities) == 0 || editEntities[0].CreateEntities == nil {
		return nil
	}
	createEntities := *editEntities[0].CreateEntities

	// create
	tx, err := handle.updateBooksWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := tx.CreateBooks(ctx, createEntities.Books); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
