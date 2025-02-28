package customizeginfx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	customizeginmiddlewares "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/middlewares"
	customizeginserver "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server"
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
			customizeginmiddlewares.NewTraceIDMiddleware,
			fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupMiddlewares.ToString())),
		),
		customizeginserver.NewGinServer,
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

type RegisterHooksParams struct {
	fx.In

	Lc          fx.Lifecycle
	Server      customizegin.GinServer
	Middlewares []customizegin.GinMiddleware `group:"middlewares"`
	Endpoints   []customizegin.Endpoint      `group:"endpoints"`
}

func registerHooks(
	params RegisterHooksParams,
) {
	params.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// register swagger
				params.Server.RegistSwagger("swagger")

				// add middlewares
				params.Server.AddMiddlewares(params.Middlewares...)

				// register endpoints
				params.Server.RegistEndPoints(params.Server.GetConfig().GetBasePath(), params.Endpoints...)

				// run server
				if err := params.Server.Run(); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("[%s] (GinServer.RunHttpServer) error in running server: {%v}", params.Server.GetConfig().GetServiceName(), err)
				}
			}()
			log.Printf("[%s] GinServer is listening on:{%s}", params.Server.GetConfig().GetServiceName(), params.Server.GetConfig().GetConnWebPort())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := params.Server.GracefulShutdown(shutdownCtx); err != nil {
				log.Printf("[%s] (GinServer.GracefulShutdown) error in shutting down server: {%v}", params.Server.GetConfig().GetServiceName(), err)
			}
			log.Printf("[%s] GinServer shutdown gracefully", params.Server.GetConfig().GetServiceName())
			log.Printf("[%s] Server shutdown", params.Server.GetConfig().GetServiceName())
			return nil
		},
	})
}
