package repositoriesredis

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/redis/go-redis/v9"
)

func NewReadBooksRedis(
	db *redis.Client,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadBooksRedis {
	return &ReadBooksRedisImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type ReadBooksRedisImpl struct {
	db            *redis.Client
	contextLogger contextloggers.ContextLogger
}

func (impl *ReadBooksRedisImpl) ReadBook(ctx context.Context, id string) (*aggregates.Book, error) {
	if id == "" {
		return nil, nil
	}
	var book documents.Book
	if err := impl.db.Get(ctx, id).Scan(&book); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read book", err)
		return nil, err
	}
	bookDocument := aggregates.BookDocument(book)
	return bookDocument.ToAggregate(), nil
}

func NewSetBookRedis(
	db *redis.Client,
	contextLogger contextloggers.ContextLogger,
) repositories.SetBookRedis {
	return &setBookRedisImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type setBookRedisImpl struct {
	db            *redis.Client
	contextLogger contextloggers.ContextLogger
}

func (impl *setBookRedisImpl) SetBook(ctx context.Context, book *aggregates.Book, timeOut time.Duration) error {
	if book == nil || book.ID == "" {
		return nil
	}
	bookDocument := book.ToDocument()
	if _, err := impl.db.Set(ctx, book.ID, *bookDocument, timeOut).Result(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to set book", err)
		return err
	}
	return nil
}

func (impl *setBookRedisImpl) DeleteBook(ctx context.Context, book *aggregates.Book) error {
	if book == nil || book.ID == "" {
		return nil
	}
	if err := impl.db.Del(ctx, book.ID).Err(); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete book", err)
		return err
	}
	return nil
}
