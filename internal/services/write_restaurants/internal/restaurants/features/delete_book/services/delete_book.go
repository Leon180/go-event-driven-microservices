package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_book/dtos"
	deleteBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_book/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type DeleteBook interface {
	DeleteBook(ctx context.Context, req *featuresdtos.DeleteBookRequest) error
}

func NewDeleteBook(
	updateBooksRepository repositories.UpdateBooks,
	readBooksRepository repositories.ReadBooks,
	rabbitmqProducer producer.Producer,
	deleteBookMessageBuilder deleteBookEvents.DeleteBookMessageBuilder,
) DeleteBook {
	return &deleteBookImpl{
		updateBooksRepository:    updateBooksRepository,
		readBooksRepository:      readBooksRepository,
		rabbitmqProducer:         rabbitmqProducer,
		deleteBookMessageBuilder: deleteBookMessageBuilder,
	}
}

type deleteBookImpl struct {
	updateBooksRepository    repositories.UpdateBooks
	readBooksRepository      repositories.ReadBooks
	rabbitmqProducer         producer.Producer
	deleteBookMessageBuilder deleteBookEvents.DeleteBookMessageBuilder
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
	if err := handle.updateBooksRepository.UpdateBook(ctx, &updateBook); err != nil {
		return err
	}

	// publish delete book event
	be := aggregates.BookEntity(*book)
	message := handle.deleteBookMessageBuilder.Build(be.ToAggregate())
	if err := handle.rabbitmqProducer.PublishMessage(ctx, message, nil, nil); err != nil {
		return err
	}

	return nil
}
