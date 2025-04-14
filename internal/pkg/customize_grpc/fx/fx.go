package customizegrpcfx

import (
	"context"
	"fmt"
	"log"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	customizegrpcinterceptors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc/interceptors"
	customizegrpcserver "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc/server"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"go.uber.org/fx"
)

// customze grpc provide module provide default grpc server:
// - grpcserver.GRPCServer
// dependencies:
// 1. server:
// - uuid.UUIDGenerator
// - loggers.Logger
// 2. server:
// - configs.GRPCConfig
var ProvideModule = fx.Module(
	"grpcserverProvideFx",
	fx.Provide(
		fx.Annotate(
			customizegrpcinterceptors.NewTraceIDUnaryInterceptor,
			fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupGRPCUnaryInterceptors.ToString())),
		),
		customizegrpcserver.NewGRPCServer,
	),
)

// customze grpc invoke module provide register hooks for grpc server
// dependencies:
// - customizegrpc.GRPCServer
// - customizegrpc.GRPCUnaryInterceptor
// - customizegrpc.GRPCServiceRegister
var InvokeModule = fx.Module(
	"grpcserverInvokeFx",
	fx.Invoke(registerHooks),
)

type RegisterHooksParams struct {
	fx.In

	Lc               fx.Lifecycle
	Server           customizegrpc.GRPCServer
	Interceptors     []customizegrpc.GRPCUnaryInterceptor `group:"grpcserverinterceptors"`
	ServiceRegisters []customizegrpc.GRPCServiceRegister  `group:"grpcserverregister"`
}

func registerHooks(
	params RegisterHooksParams,
) {
	params.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// add interceptors
				params.Server.AddUnaryInterceptors(params.Interceptors...)

				// register services
				params.Server.RegistGRPCService(params.ServiceRegisters...)

				// run server
				if err := params.Server.Run(); err != nil {
					log.Fatalf(
						"[%s] (GRPCServer.Run) error in running server: {%v}",
						params.Server.GetConfig().GetServiceName(),
						err,
					)
				}
			}()
			log.Printf(
				"[%s] GRPCServer is listening on:{%s}",
				params.Server.GetConfig().GetServiceName(),
				params.Server.GetConfig().GetPort(),
			)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Server.GracefulShutdown()
			log.Printf("[%s] GRPCServer shutdown gracefully", params.Server.GetConfig().GetServiceName())
			return nil
		},
	})
}
