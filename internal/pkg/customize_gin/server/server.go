package customizeginserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ginconfigs "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/configs"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	customizeginmiddlewares "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/middlewares"
	zaplogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/zap_logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/location"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type GinServer interface {
	RegistSwagger(swaggerPath string)
	RegistEndPoints(routerGroupName string, endpoints ...customizeginendpoints.Endpoint)
	AddMiddlewares(middlewares ...customizeginmiddlewares.GinMiddleware)

	Run() error
	GracefulShutdown(ctx context.Context) error
	GetConfig() ginconfigs.GinConfig
}

func NewGinServer(
	cfg ginconfigs.GinConfig,
	zapLogger zaplogger.ZapLogger,
) GinServer {
	engine := gin.Default()
	if err := engine.SetTrustedProxies(nil); err != nil {
		zapLogger.Fatal("error while setting trusted proxies", err)
	}
	engine.Use(
		ginzap.Ginzap(zapLogger.InternalLogger(), time.RFC3339, true),
		ginzap.RecoveryWithZap(zapLogger.InternalLogger(), true),
		gzip.Gzip(gzip.DefaultCompression),
		location.Default(),
		cors.New(cfg.GenerateCORSConfig()),
	)
	return &GinServerImpl{
		engine: engine,
		config: cfg,
	}
}

type GinServerImpl struct {
	engine *gin.Engine
	config ginconfigs.GinConfig
	server *http.Server
}

func (s *GinServerImpl) AddMiddlewares(middlewares ...customizeginmiddlewares.GinMiddleware) {
	for _, middleware := range middlewares {
		s.engine.Use(middleware.Handle())
	}
}

func (s *GinServerImpl) RegistSwagger(swaggerPath string) {
	// Disable compression for Swagger endpoints
	noCompression := s.engine.Group(swaggerPath)
	{
		noCompression.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.URL(fmt.Sprintf("%s/doc.json", swaggerPath)), // The URL pointing to API definition
			ginSwagger.DefaultModelsExpandDepth(-1),
		))
	}
}

func (s *GinServerImpl) RegistEndPoints(routerGroupName string, endpoints ...customizeginendpoints.Endpoint) {
	routerGroup := s.engine.Group(s.engine.BasePath()).Group(routerGroupName)
	for _, endpoint := range endpoints {
		endpoint.MapEndpoint(routerGroup)
	}
}

func (s *GinServerImpl) Run() error {
	s.server = &http.Server{
		Addr:    ":" + s.config.GetConnWebPort(),
		Handler: s.engine,
	}
	return s.server.ListenAndServe()
}

func (s *GinServerImpl) GracefulShutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *GinServerImpl) GetConfig() ginconfigs.GinConfig {
	return s.config
}
