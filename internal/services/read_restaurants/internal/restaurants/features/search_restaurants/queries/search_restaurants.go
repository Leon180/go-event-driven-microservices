package queries

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	utilitiesdb "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/db"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type SearchRestaurants struct {
	NameFilter        *string
	NamePreciseSearch bool

	DescriptionFilter       *string
	CityFilter              []enums.City
	CountryFilter           []enums.Country
	MaxPriceFilter          *int
	MinPriceFilter          *int
	CategoryFilter          []enums.Category
	TableAvailableWeek      []time.Weekday
	TableAvailableStartTime *string
	TableAvailableEndTime   *string
	OrderBy                 []utilitiesdb.OrderBy
	Pagination              *utilitiesdb.Pagination
}

type SearchRestaurantsHandler interface {
	SearchRestaurants(
		ctx context.Context,
		command *SearchRestaurants,
	) (aggregates.Restaurants, error)
}

func NewSearchRestaurantsHandler(
	searchRestaurantsRepository repositories.SearchRestaurantsMongo,
) SearchRestaurantsHandler {
	return &searchRestaurantsImpl{searchRestaurantsRepository: searchRestaurantsRepository}
}

type searchRestaurantsImpl struct {
	searchRestaurantsRepository repositories.SearchRestaurantsMongo
}

func (handle *searchRestaurantsImpl) SearchRestaurants(
	ctx context.Context,
	command *SearchRestaurants,
) (aggregates.Restaurants, error) {
	if command == nil {
		return nil, nil
	}
	dtos := dtos.SearchRestaurants{
		NameFilter:              command.NameFilter,
		NamePreciseSearch:       command.NamePreciseSearch,
		DescriptionFilter:       command.DescriptionFilter,
		CityFilter:              command.CityFilter,
		CountryFilter:           command.CountryFilter,
		MaxPriceFilter:          command.MaxPriceFilter,
		MinPriceFilter:          command.MinPriceFilter,
		CategoryFilter:          command.CategoryFilter,
		TableAvailableWeek:      command.TableAvailableWeek,
		TableAvailableStartTime: command.TableAvailableStartTime,
		TableAvailableEndTime:   command.TableAvailableEndTime,
		OrderBy:                 command.OrderBy,
		Pagination:              command.Pagination,
	}
	return handle.searchRestaurantsRepository.SearchRestaurants(ctx, &dtos)
}
