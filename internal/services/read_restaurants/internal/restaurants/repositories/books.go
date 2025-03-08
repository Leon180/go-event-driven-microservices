package repositories

import (
	"context"

	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
)

//go:generate mockgen -source=books.go -destination=./mocks/books_mock.go -package=mocks

type SearchBooksFullInfo interface {
	SearchBooksFullInfo(ctx context.Context, searchBooks *dtos.SearchBooks) (aggregates.Books, error)
}

type ReadBooks interface {
	ReadBookFullInfo(ctx context.Context, id string) (*aggregates.Book, error)
	ReadBook(ctx context.Context, id string) (*entities.Book, error)
}

type UpdateBooks interface {
	CreateBooks(ctx context.Context, books entities.Books) error
	UpdateBook(ctx context.Context, updateBook *entities.UpdateBook) error
	DeleteBooks(ctx context.Context, ids []string) error
}

type UpdateBooksWithTransaction interface {
	customizegorm.Transaction
	UpdateBooks
}
