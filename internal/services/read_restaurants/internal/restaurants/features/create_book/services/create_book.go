package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type CreateBookHandler interface {
	CreateBook(ctx context.Context, aggregate *aggregates.Book) error
}

func NewCreateBookHandler(
	redisConfig *redisdb.RedisConfig,
	updateBooksMongo repositories.UpdateBooksMongo,
	readBooksMongo repositories.ReadBooksMongo,
	setBookRedis repositories.SetBookRedis,
) CreateBookHandler {
	return &createBookImpl{
		redisConfig:      redisConfig,
		updateBooksMongo: updateBooksMongo,
		readBooksMongo:   readBooksMongo,
		setBookRedis:     setBookRedis,
	}
}

type createBookImpl struct {
	redisConfig      *redisdb.RedisConfig
	updateBooksMongo repositories.UpdateBooksMongo
	readBooksMongo   repositories.ReadBooksMongo
	setBookRedis     repositories.SetBookRedis
}

func (handle *createBookImpl) CreateBook(ctx context.Context, aggregate *aggregates.Book) error {
	if aggregate == nil {
		return nil
	}

	// check if book already exists
	book, err := handle.readBooksMongo.ReadBook(ctx, aggregate.ID)
	if err != nil {
		return err
	}
	if book != nil {
		if book.ActiveStatus {
			return customizeerrors.BookAlreadyExistsError
		}
		return customizeerrors.BookAlreadyExistsButInactiveError
	}

	err = handle.updateBooksMongo.CreateBooks(ctx, []aggregates.Book{*aggregate})
	if err != nil {
		return err
	}

	err = handle.setBookRedis.SetBook(ctx, aggregate, time.Duration(handle.redisConfig.CacheTimeOut)*time.Second)
	if err != nil {
		return err
	}

	return nil
}
