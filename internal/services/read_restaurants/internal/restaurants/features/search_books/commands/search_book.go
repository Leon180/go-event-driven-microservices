package commands

import (
	"context"
	"time"

	utilitiesdb "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/db"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type SearchBooks struct {
	MobileNumber *string
	TableID      *string
	AvailableID  *string

	NameFilter        *string
	NamePreciseSearch bool

	TableAvailableWeek      []time.Weekday
	TableAvailableStartTime *string
	TableAvailableEndTime   *string

	OrderBy    []utilitiesdb.OrderBy
	Pagination *utilitiesdb.Pagination
}

type SearchBooksHandler interface {
	SearchBooks(
		ctx context.Context,
		command *SearchBooks,
	) (aggregates.Books, error)
}

func NewSearchBooksHandler(
	searchBooksRepository repositories.SearchBooksMongo,
) SearchBooksHandler {
	return &searchBooksImpl{searchBooksRepository: searchBooksRepository}
}

type searchBooksImpl struct {
	searchBooksRepository repositories.SearchBooksMongo
}

func (handle *searchBooksImpl) SearchBooks(
	ctx context.Context,
	command *SearchBooks,
) (aggregates.Books, error) {
	if command == nil {
		return nil, nil
	}
	dtos := dtos.SearchBooks{
		MobileNumber:            command.MobileNumber,
		TableID:                 command.TableID,
		AvailableID:             command.AvailableID,
		NameFilter:              command.NameFilter,
		NamePreciseSearch:       command.NamePreciseSearch,
		TableAvailableWeek:      command.TableAvailableWeek,
		TableAvailableStartTime: command.TableAvailableStartTime,
		TableAvailableEndTime:   command.TableAvailableEndTime,
		OrderBy:                 command.OrderBy,
		Pagination:              command.Pagination,
	}
	return handle.searchBooksRepository.SearchBooks(ctx, &dtos)
}
