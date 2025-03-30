package events

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
)

type CreateBook struct {
	*types.MessageImpl
	aggregates.Book
}

type CreateBookMessageBuilder interface {
	Build(book *aggregates.Book) *CreateBook
}

func NewCreateBookMessageBuilder(uuidGenerator uuid.UUIDGenerator) CreateBookMessageBuilder {
	return &createBookMessageBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type createBookMessageBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *createBookMessageBuilderImpl) Build(book *aggregates.Book) *CreateBook {
	return &CreateBook{
		Book:        *book,
		MessageImpl: types.NewMessageImpl(b.uuidGenerator.GenerateUUID(), "CreateBook"),
	}
}
