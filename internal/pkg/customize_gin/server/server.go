package customizeginserver

import (
	"context"

	ginconfigs "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/configs"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	customizeginmiddlewares "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/middlewares"
)

type GinServer interface {
	RegistSwagger(swaggerPath string)
	RegistEndPoints(routerGroupName string, endpoints ...customizeginendpoints.Endpoint)
	AddMiddlewares(middlewares ...customizeginmiddlewares.GinMiddleware)

	Run() error
	GracefulShutdown(ctx context.Context) error
	GetConfig() ginconfigs.GinConfig
}
