package integrationevents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
)

type DeleteBook struct {
	types.Message
	aggregates.Book
}

type DeleteBookBuilder interface {
	Build(book aggregates.Book) *DeleteBook
}

func NewDeleteBookBuilder(uuidGenerator uuid.UUIDGenerator) DeleteBookBuilder {
	return &deleteBookBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type deleteBookBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *deleteBookBuilderImpl) Build(book aggregates.Book) *DeleteBook {
	return &DeleteBook{
		Book:    book,
		Message: types.NewMessage(b.uuidGenerator.GenerateUUID(), "DeleteBook"),
	}
}
