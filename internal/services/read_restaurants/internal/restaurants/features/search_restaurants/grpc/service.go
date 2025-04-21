package grpc

import (
	"context"
	"time"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	utilitiesdb "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/db"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	aggregatesconvert "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/convert/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/protobuf"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/queries"
	"github.com/samber/lo"
	"google.golang.org/grpc"
)

func NewGRPCServiceRegister(
	searchRestaurantsQuery queries.SearchRestaurantsHandler,
) customizegrpc.GRPCServiceRegister {
	return func(s *grpc.Server) {
		protobuf.RegisterSearchRestaurantsServiceServer(s, NewGRPCService(searchRestaurantsQuery))
	}
}

func NewGRPCService(
	searchRestaurantsQuery queries.SearchRestaurantsHandler,
) *grpcService {
	return &grpcService{
		searchRestaurantsQuery: searchRestaurantsQuery,
	}
}

type grpcService struct {
	protobuf.UnimplementedSearchRestaurantsServiceServer
	searchRestaurantsQuery queries.SearchRestaurantsHandler
}

func (s *grpcService) SearchRestaurants(
	ctx context.Context,
	req *protobuf.SearchRestaurantsReq,
) (*protobuf.Restaurants, error) {
	if req == nil {
		return nil, nil
	}
	restaurants, err := s.searchRestaurantsQuery.SearchRestaurants(ctx, &queries.SearchRestaurants{
		NameFilter:        req.NameFilter,
		NamePreciseSearch: req.NamePreciseSearch,
		BranchID:          req.BranchID,
		DescriptionFilter: req.DescriptionFilter,
		CityFilter:        lo.Map(req.CityFilter, func(city string, _ int) enums.City { return enums.City(city) }),
		CountryFilter: lo.Map(
			req.CountryFilter,
			func(country string, _ int) enums.Country { return enums.Country(country) },
		),
		MaxPriceFilter: func() *int {
			if req.MaxPriceFilter != nil {
				maxPrice := int(*req.MaxPriceFilter)
				return &maxPrice
			}
			return nil
		}(),
		MinPriceFilter: func() *int {
			if req.MinPriceFilter != nil {
				minPrice := int(*req.MinPriceFilter)
				return &minPrice
			}
			return nil
		}(),
		CategoryFilter: lo.Map(
			req.CategoryFilter,
			func(category string, _ int) enums.Category { return enums.Category(category) },
		),
		TableAvailableWeek: lo.Map(
			req.TableAvailableWeek,
			func(week int32, _ int) time.Weekday { return time.Weekday(week) },
		),
		TableAvailableStartTime: req.TableAvailableStartTime,
		TableAvailableEndTime:   req.TableAvailableEndTime,
		OrderBy: lo.Map(req.OrderBy, func(orderBy *protobuf.OrderBy, _ int) utilitiesdb.OrderBy {
			return utilitiesdb.OrderBy{
				Field:     orderBy.Field,
				Direction: utilitiesdb.OrderByDirection(orderBy.Direction),
			}
		}),
		Pagination: func() *utilitiesdb.Pagination {
			if req.Pagination != nil {
				return &utilitiesdb.Pagination{
					Page:     int(req.Pagination.Page),
					PageSize: int(req.Pagination.PageSize),
				}
			}
			return nil
		}(),
	})
	if err != nil {
		return nil, err
	}
	return &protobuf.Restaurants{
		Restaurants: lo.Map(restaurants, func(restaurant aggregates.Restaurant, _ int) *protobuf.Restaurant {
			ar := aggregatesconvert.RestaurantAggregate(restaurant)
			return ar.ToProto()
		}),
	}, nil
}
