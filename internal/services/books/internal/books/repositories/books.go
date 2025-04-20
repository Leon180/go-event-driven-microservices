package repositories

import (
	"context"

	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
)

//go:generate mockgen -source=books.go -destination=./mocks/books_mock.go -package=mocks

type SearchBooksFullInfo interface {
	SearchBooksFullInfo(ctx context.Context, searchBooks *dtos.SearchBooks) (entities.Books, error)
}

type ReadBook interface {
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
