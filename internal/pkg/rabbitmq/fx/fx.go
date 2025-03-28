package rabbitmqfx

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	bus "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/bus"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	rabbitmqbus "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/bus"
	rabbitmqconnect "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/connect"
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqproducer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/producer"

	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"rabbitmqfx",
	fx.Provide(rabbitmq.NewRabbitMQConfig),
	fx.Provide(rabbitmqconnect.NewRabbitMQConnection),
	fx.Provide(rabbitmqconsumer.NewConsumerFactory),
	fx.Provide(rabbitmqproducer.NewProducerFactory),
	fx.Provide(fx.Annotate(
		rabbitmqbus.NewRabbitMQBus,
		fx.ParamTags(``, ``, ``, `optional:"true"`),
		fx.As(new(producer.Producer)),
		fx.As(new(bus.Bus)),
		fx.As(new(rabbitmqbus.RabbitMQBus)),
	)),
)

var InvokeModule = fx.Options(
	fx.Invoke(registerHooks),
)

// we don't want to register any dependencies here, its func body should execute always even we don't request for that, so we should use `invoke`
func registerHooks(
	lc fx.Lifecycle,
	bus rabbitmqbus.RabbitMQBus,
	rabbitmqOptions *rabbitmq.RabbitMQConfig,
	logger loggers.Logger,
) {
	if !rabbitmqOptions.AutoStart {
		return
	}

	lifeTimeCtx := context.Background()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := bus.Start(lifeTimeCtx); err != nil {
					logger.Errorf("(bus.Start) error in running rabbitmq server: {%v}", err)
					return
				}
			}()
			logger.Info("rabbitmq is listening.")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := bus.Stop(); err != nil {
				logger.Errorf("error shutting down rabbitmq server: %v", err)
			} else {
				logger.Info("rabbitmq server shutdown gracefully")
			}
			_, cancel := context.WithTimeout(lifeTimeCtx, 5*time.Second)
			defer cancel()
			return nil
		},
	})
}
