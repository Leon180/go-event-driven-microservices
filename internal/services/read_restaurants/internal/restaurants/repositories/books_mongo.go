package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
)

//go:generate mockgen -source=books.go -destination=./mocks/books_mock.go -package=mocks

type SearchBooksMongo interface {
	SearchBooks(ctx context.Context, searchBooks *dtos.SearchBooks) (aggregates.Books, error)
}

type ReadBooksMongo interface {
	ReadBook(ctx context.Context, id string) (*aggregates.Book, error)
}

type UpdateBooksMongo interface {
	CreateBooks(ctx context.Context, books aggregates.Books) error
	UpdateBook(ctx context.Context, updateBook *aggregates.Book) error
	DeleteBooks(ctx context.Context, ids []string) error
}
