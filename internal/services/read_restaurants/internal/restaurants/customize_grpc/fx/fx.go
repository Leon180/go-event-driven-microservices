package grpcfx

import (
	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"restaurantsProvideGRPCFx",

	fx.Provide(
		customizegrpc.NewGRPCConfig,
	),
)
