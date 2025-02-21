package customizeginfx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	ginserver "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server"
	ginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	ginmiddlewares "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/middlewares"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"go.uber.org/fx"
)

// customze gin provide module provide default gin middlewares and gin server:
// - ginmiddlewares.GinMiddleware(ginmiddlewares.TraceIDMiddleware)
// - ginserver.GinServer
// dependencies:
// 1. middlewares:
// - uuid.UUIDGenerator
// - loggers.Logger
// 2. server:
// - zaplogger.Logger
// - configs.GinConfig
var ProvideModule = fx.Module(
	"ginserverProvideFx",
	fx.Provide(
		fx.Annotate(
			ginmiddlewares.NewTraceIDMiddleware,
			fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupMiddlewares.ToString())),
		),
		ginserver.NewGinServer,
	),
)

// customze gin invoke module provide register hooks for gin server and endpoints
// dependencies:
// - server.GinServer
// - middlewares.GinMiddleware
// - endpoints.Endpoint
var InvokeModule = fx.Module(
	"ginserverInvokeFx",
	fx.Invoke(registerHooks),
)

type registerHooksParams struct {
	fx.In

	lc          fx.Lifecycle
	server      ginserver.GinServer
	middlewares []ginmiddlewares.GinMiddleware `group:"middlewares"`
	endpoints   []ginendpoints.Endpoint        `group:"endpoints"`
}

func registerHooks(
	params registerHooksParams,
) {
	params.lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// add middlewares
				params.server.AddMiddlewares(params.middlewares...)

				// register endpoints
				params.server.RegistEndPoints(params.server.GetConfig().GetServiceName(), params.endpoints...)

				// run server
				if err := params.server.Run(); !errors.Is(
					err,
					http.ErrServerClosed,
				) {
					// do a fatal for going to OnStop process
					log.Fatalf(
						"[%s] (GinServer.RunHttpServer) error in running server: {%v}",
						params.server.GetConfig().GetServiceName(),
						err,
					)
				}
			}()
			log.Printf("[%s] GinServer is listening on:{%s}", params.server.GetConfig().GetServiceName(), params.server.GetConfig().GetConnWebPort())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := params.server.GracefulShutdown(shutdownCtx); err != nil {
				log.Printf("[%s] (GinServer.GracefulShutdown) error in shutting down server: {%v}", params.server.GetConfig().GetServiceName(), err)
			}
			log.Printf("[%s] GinServer shutdown gracefully", params.server.GetConfig().GetServiceName())
			log.Printf("[%s] Server shutdown", params.server.GetConfig().GetServiceName())
			return nil
		},
	})
}
