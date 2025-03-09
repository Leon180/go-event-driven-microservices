package repositories

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
)

//go:generate mockgen -source=books_redis.go -destination=./mocks/books_redis_mock.go -package=mocks

type ReadBooksRedis interface {
	ReadBook(ctx context.Context, id string) (*aggregates.Book, error)
}

type SetBookRedis interface {
	SetBook(ctx context.Context, book *aggregates.Book, timeOut time.Duration) error
}
