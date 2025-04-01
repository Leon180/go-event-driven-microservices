package repositoriesredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
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
	var data []byte
	if err := impl.db.Get(ctx, id).Scan(&data); err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read book", err)
		return nil, err
	}
	var book aggregates.Book
	if err := json.Unmarshal(data, &book); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to unmarshal book when reading from redis: ", err)
		return nil, err
	}
	return &book, nil
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
	data, err := json.Marshal(*book)
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to marshal book when setting to redis: ", err)
		return err
	}
	if _, err := impl.db.Set(ctx, book.ID, data, timeOut).Result(); err != nil {
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
		if err == redis.Nil {
			return nil
		}
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete book", err)
		return err
	}
	return nil
}
