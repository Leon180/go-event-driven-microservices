package integrationevents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
)

type CreateBook struct {
	types.Message
	aggregates.Book
}

type CreateBookBuilder interface {
	Build(book aggregates.Book) *CreateBook
}

func NewCreateBookBuilder(uuidGenerator uuid.UUIDGenerator) CreateBookBuilder {
	return &createBookBuilderImpl{
		uuidGenerator: uuidGenerator,
	}
}

type createBookBuilderImpl struct {
	uuidGenerator uuid.UUIDGenerator
}

func (b *createBookBuilderImpl) Build(book aggregates.Book) *CreateBook {
	return &CreateBook{
		Book:    book,
		Message: types.NewMessage(b.uuidGenerator.GenerateUUID(), "CreateBook"),
	}
}
