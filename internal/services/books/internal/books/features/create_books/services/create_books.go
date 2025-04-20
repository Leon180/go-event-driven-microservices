package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/aggregates"
	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories"
	"github.com/samber/lo"
)

type CreateBooks interface {
	CreateBooks(ctx context.Context, req *featuresdtos.CreateBooksRequest) error
}

func NewCreateBooks(
	uuidGenerator uuid.UUIDGenerator,
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo,
	grpcBookService customizegrpc.GRPCBookService,
) CreateBooks {
	return &createBooksImpl{
		uuidGenerator:                        uuidGenerator,
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		searchBooksFullInfoRepository:        searchBooksFullInfoRepository,
		grpcBookService:                      grpcBookService,
	}
}

type createBooksImpl struct {
	uuidGenerator                        uuid.UUIDGenerator
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	searchBooksFullInfoRepository        repositories.SearchBooksFullInfo
	grpcBookService                      customizegrpc.GRPCBookService
}

func (handle *createBooksImpl) CreateBooks(ctx context.Context, req *featuresdtos.CreateBooksRequest) error {
	if req == nil {
		return nil
	}
	// check if book already exists
	books, err := handle.searchBooksFullInfoRepository.SearchBooksFullInfo(ctx, &dtos.SearchBooks{
		RestaurantID: &req.RestaurantID,
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
	if err := bookDTOAggregateBuilder.SaveBook(req); err != nil {
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

	// publish create book events
	aggregates := bookDTOAggregateBuilder.GetAggregates()
	for _, book := range aggregates {
		message := handle.createBookMessageBuilder.Build(&book)
		if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
			return err
		}
	}

	return nil
}
