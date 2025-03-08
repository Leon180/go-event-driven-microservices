package redisdbfx

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/redisdb"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

// ProvideModule is the module for the redis
// It provides:
// - *redis.RedisConfig
// - *redis.RedisClient
// - redis.UniversalClient
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"redisProvideFx",
	fx.Provide(
		redisdb.NewRedisConfig,
		redisdb.NewRedisClient,
		func(client *redis.Client) redis.UniversalClient {
			return client
		},
	),
)

// InvokeModule is the module for the redis
// It invokes the registerHooks function
// dependencies:
// - redis.UniversalClient
// - logger.Logger
var InvokeModule = fx.Module(
	"redisInvokeFx",
	fx.Invoke(
		registerHooks,
	),
)

func registerHooks(
	lc fx.Lifecycle,
	client redis.UniversalClient,
	logger loggers.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := client.Ping(ctx).Err(); err != nil {
				logger.Errorf("error in pinging redis: %v", err)
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := client.Close(); err != nil {
				logger.Errorf("error in closing redis: %v", err)
			}
			return nil
		},
	})
}
