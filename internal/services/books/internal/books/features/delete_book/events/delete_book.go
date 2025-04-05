package integrationevents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/aggregates"
)

type DeleteBook struct {
	*types.MessageImpl
	aggregates.Book
}

type DeleteBookMessageBuilder interface {
	Build(book *aggregates.Book) *DeleteBook
}

func NewDeleteBookMessageBuilder(uuidGenerator uuid.UUIDGenerator) DeleteBookMessageBuilder {
	return &deleteBookMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type deleteBookMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *deleteBookMessageBuilderImpl) Build(book *aggregates.Book) *DeleteBook {
	return &DeleteBook{
		Book:        *book,
		MessageImpl: types.NewMessageImpl(b.uuidGenerator.GenerateUUID(), "DeleteBook"),
	}
}
