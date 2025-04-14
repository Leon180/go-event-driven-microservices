package grpc

import (
	"context"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	aggregatesconvert "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/convert/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/protobuf"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/queries"
	"google.golang.org/grpc"
)

func NewGRPCServiceRegister(
	getRestaurantQuery queries.GetRestaurantHandler,
) customizegrpc.GRPCServiceRegister {
	return func(s *grpc.Server) {
		protobuf.RegisterGetRestaurantServiceServer(s, NewGRPCService(getRestaurantQuery))
	}
}

func NewGRPCService(
	getRestaurantQuery queries.GetRestaurantHandler,
) *grpcService {
	return &grpcService{
		getRestaurantQuery: getRestaurantQuery,
	}
}

type grpcService struct {
	protobuf.UnimplementedGetRestaurantServiceServer
	getRestaurantQuery queries.GetRestaurantHandler
}

func (s *grpcService) GetRestaurant(ctx context.Context, req *protobuf.GetRestaurantReq) (*protobuf.Restaurant, error) {
	if req == nil {
		return nil, nil
	}
	if req.RestaurantId == "" {
		return nil, nil
	}
	restaurant, err := s.getRestaurantQuery.GetRestaurant(ctx, &queries.GetRestaurant{ID: req.RestaurantId})
	if err != nil {
		return nil, err
	}
	if restaurant == nil {
		return nil, nil
	}
	ra := aggregatesconvert.RestaurantAggregate(*restaurant)
	return ra.ToProto(), nil
}
