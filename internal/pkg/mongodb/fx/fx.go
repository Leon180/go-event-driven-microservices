package mongodbfx

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the mongodb
// It provides:
// - *MongoDbConfig
// - *mongo.Client
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"mongoProvideFx",
	fx.Provide(
		mongodb.NewMongoDbConfig,
		mongodb.NewMongoDB,
	),
)

// InvokeModule is the module for the mongodb
// It invokes the registerHooks function
// dependencies:
// - *mongo.Client
// - logger.Logger
var InvokeModule = fx.Module(
	"mongoInvokeFx",
	fx.Invoke(
		registerHooks,
	),
)

func registerHooks(
	lc fx.Lifecycle,
	client *mongo.Client,
	logger loggers.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := client.Ping(ctx, nil)
			if err != nil {
				logger.Error("failed to ping mongo", err)
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := client.Disconnect(ctx); err != nil {
				logger.Errorf("error in disconnecting mongo: %v", err)
			}
			return nil
		},
	})
}
