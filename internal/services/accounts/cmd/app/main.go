package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/middlewares"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	accountconfigs "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/location"
	ginzap "github.com/gin-contrib/zap"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg           configs.Config
		accountConfig accountconfigs.Config
		db            *gorm.DB
		engine        *gin.Engine
	)
	if err := configs.LoadConfig(&cfg, "../../../../../config"); err != nil {
		log.Fatalf("cannot load config: %+v", err)
	}
	if err := accountconfigs.LoadConfig(&accountConfig, "../../../config"); err != nil {
		log.Fatalf("cannot load config: %+v", err)
	}
	logger, _ := utilities.InitLogger(cfg.GenLogConfig())
	defer utilities.SyncLogger()

	// init postgresql db
	db = postgresdb.InitGormDB(accountConfig, logger)
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			utilities.LogError(context.Background(), err)
		}
		if err := sqlDB.Close(); err != nil {
			utilities.LogError(context.Background(), err)
		}
	}()
	if err := migrations.MigrateDB(db); err != nil {
		utilities.LogError(context.Background(), err)
	}

	// init middleware
	traceIDMiddleware := middlewares.NewTraceIDMiddleware(utilities.NewUUIDGenerator())

	engine = gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(
		ginzap.Ginzap(logger, time.RFC3339, true),
		ginzap.RecoveryWithZap(logger, true),
		gzip.Gzip(gzip.DefaultCompression),
		location.Default(),
		cors.New(cfg.GenCORSConfig()),
		traceIDMiddleware.Handler(),
	)
	// setRoute(engine, controllerHandle)
	server := &http.Server{
		Addr:    ":" + cfg.ConnWebPort,
		Handler: engine,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			utilities.LogFatal(context.Background(), err)
		}
	}()

	// lis, err := net.Listen("tcp", ":"+cfg.ConnGRPCPort) // Add cfg.GRPCPort to your config
	// if err != nil {
	// 	utility.LogFatal(context.Background(), "failed to listen: %v", err)
	// }
	// traceIDInterceptor := middleware.NewTraceIDInterceptor()
	// s := grpc.NewServer(
	// 	grpc.UnaryInterceptor(traceIDInterceptor.Unary()),
	// )
	// setGRPCService(s, controllerHandle)
	// reflection.Register(s)
	// go func() {
	// 	if err := s.Serve(lis); err != nil {
	// 		utility.SugarLogger.Fatal(err)
	// 	}
	// }()

	// graceful shutdown setup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// graceful shutdown
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := server.Shutdown(shutdownCtx); err != nil {
		utilities.LogError(context.Background(), "Error during server shutdown: %v", err)
	}
	// s.GracefulStop()
	utilities.LogInfo(context.Background(), "[ACCOUNTS-SERVICE] Shutting down gracefully")
	utilities.LogInfo(context.Background(), "[ACCOUNTS-SERVICE] Server shutdown")
}

// func setRoute(engine *gin.Engine, controllerHandle *inject.ControllerHandle) {
// 	// Swagger docs
// 	// Use this URL config for swagger
// 	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
// 	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
// 	defaultRouter := engine.Group(engine.BasePath())
// 	baseRouter := defaultRouter.Group("/tabelogo-spider")
// 	baseRouter.GET("/getTabelogInfo", controllerHandle.GetTabelogInfoController.GetTabelogInfo)
// 	baseRouter.GET("/getTabelogPhoto", controllerHandle.GetTabelogInfoController.GetTabelogPhoto)
// }

// func setGRPCService(s *grpc.Server, controllerHandle *inject.ControllerHandle) {
// 	proto.RegisterTabelogoSpiderServiceServer(s, controllerHandle.TabelogoSpiderServiceServer)
// }
