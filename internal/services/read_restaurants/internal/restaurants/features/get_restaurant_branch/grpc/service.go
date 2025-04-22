package grpc

import (
	"context"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	aggregatesconvert "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/convert/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/protobuf"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant_branch/queries"
	"google.golang.org/grpc"
)

func NewGRPCServiceRegister(
	getRestaurantBranchQuery queries.GetRestaurantBranchHandler,
) customizegrpc.GRPCServiceRegister {
	return func(s *grpc.Server) {
		protobuf.RegisterGetRestaurantBranchServiceServer(s, NewGRPCService(getRestaurantBranchQuery))
	}
}

func NewGRPCService(
	getRestaurantBranchQuery queries.GetRestaurantBranchHandler,
) *grpcService {
	return &grpcService{
		getRestaurantBranchQuery: getRestaurantBranchQuery,
	}
}

type grpcService struct {
	protobuf.UnimplementedGetRestaurantBranchServiceServer
	getRestaurantBranchQuery queries.GetRestaurantBranchHandler
}

func (s *grpcService) GetRestaurantBranch(
	ctx context.Context,
	req *protobuf.GetRestaurantBranchReq,
) (*protobuf.Restaurant, error) {
	if req == nil {
		return nil, nil
	}
	if req.BranchId == "" {
		return nil, nil
	}
	restaurant, err := s.getRestaurantBranchQuery.GetRestaurantBranch(
		ctx,
		&queries.GetRestaurantBranch{BranchID: req.BranchId},
	)
	if err != nil {
		return nil, err
	}
	if restaurant == nil {
		return nil, nil
	}
	ra := aggregatesconvert.RestaurantAggregate(*restaurant)
	return ra.ToProto(), nil
}
