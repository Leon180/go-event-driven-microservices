package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_book/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type CreateBook interface {
	CreateBook(ctx context.Context, req *dtos.Book) error
}

func NewCreateBook(
	uuidGenerator uuid.UUIDGenerator,
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo,
	rabbitmqProducer producer.Producer,
	createBookMessageBuilder events.CreateBookMessageBuilder,
) CreateBook {
	return &createBookImpl{
		uuidGenerator:                        uuidGenerator,
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		searchBooksFullInfoRepository:        searchBooksFullInfoRepository,
		rabbitmqProducer:                     rabbitmqProducer,
		createBookMessageBuilder:             createBookMessageBuilder,
	}
}

type createBookImpl struct {
	uuidGenerator                        uuid.UUIDGenerator
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	searchBooksFullInfoRepository        repositories.SearchBooksFullInfo
	rabbitmqProducer                     producer.Producer
	createBookMessageBuilder             events.CreateBookMessageBuilder
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
