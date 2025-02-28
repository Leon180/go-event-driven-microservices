package customizegin

import (
	"context"
)

type GinServer interface {
	RegistSwagger(swaggerPath string)
	RegistEndPoints(routerGroupName string, endpoints ...Endpoint)
	AddMiddlewares(middlewares ...GinMiddleware)

	Run() error
	GracefulShutdown(ctx context.Context) error
	GetConfig() GinConfig
}
