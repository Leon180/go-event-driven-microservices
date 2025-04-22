package aggregates

import "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"

type Book entities.Book

func (b *Book) ToEntity() *entities.Book {
	bs := *b
	eb := entities.Book(bs)
	return &eb
}
