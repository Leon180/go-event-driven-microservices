package customizegrpc

import (
	"context"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	customizegrpcclient "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc/client"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc/protobuf"
)

type GRPCBookService interface {
	GetRestaurant(ctx context.Context, req *protobuf.GetRestaurantReq) (*protobuf.Restaurant, error)
	SearchRestaurants(ctx context.Context, req *protobuf.SearchRestaurantsReq) (*protobuf.Restaurants, error)
}

func NewGRPCBookService(
	config customizegrpc.GRPCConfig,
	logger contextloggers.ContextLogger,
) (GRPCBookService, error) {
	client, err := customizegrpcclient.NewGRPCClient(config)
	if err != nil {
		return nil, err
	}
	conn := client.GetConnection()
	getRestaurantServiceClient := protobuf.NewGetRestaurantServiceClient(conn)
	searchRestaurantsServiceClient := protobuf.NewSearchRestaurantsServiceClient(conn)

	return &bookService{
		GRPCClient:                     client,
		getRestaurantServiceClient:     getRestaurantServiceClient,
		searchRestaurantsServiceClient: searchRestaurantsServiceClient,
		logger:                         logger,
	}, nil
}

type bookService struct {
	customizegrpc.GRPCClient
	getRestaurantServiceClient     protobuf.GetRestaurantServiceClient
	searchRestaurantsServiceClient protobuf.SearchRestaurantsServiceClient
	logger                         contextloggers.ContextLogger
}

func (b *bookService) GetRestaurant(ctx context.Context, req *protobuf.GetRestaurantReq) (*protobuf.Restaurant, error) {
	rest, err := b.getRestaurantServiceClient.GetRestaurant(ctx, req)
	if err != nil {
		b.logger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("Failed to get restaurant")
		return nil, err
	}
	return rest, nil
}

func (b *bookService) SearchRestaurants(ctx context.Context, req *protobuf.SearchRestaurantsReq) (*protobuf.Restaurants, error) {
	rests, err := b.searchRestaurantsServiceClient.SearchRestaurants(ctx, req)
	if err != nil {
		b.logger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("Failed to search restaurants")
		return nil, err
	}
	return rests, nil
}
