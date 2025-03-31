package rabbitmqconsumer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewDeadLetterArgs(rabbitMQConfig *rabbitmq.RabbitMQConfig) *amqp.Table {
	return &amqp.Table{
		enums.DeliveryHeaderDeadLetterExchange.ToString():   rabbitMQConfig.DeadLetterExchange,
		enums.DeliveryHeaderDeadLetterRoutingKey.ToString(): rabbitMQConfig.DeadLetterRoutingKey,
	}
}

type DeadLetterRabbitMQConsumerConfig struct {
	RabbitMQConsumerConfig
	ConsumerMessageTypes []types.Message
	DeadLetterArgs       *amqp.Table
}

func NewDefaultDeadLetterRabbitMQConsumerConfig(
	consumerMessageTypes []types.Message,
	rabbitMQConfig *rabbitmq.RabbitMQConfig,
) *DeadLetterRabbitMQConsumerConfig {
	return &DeadLetterRabbitMQConsumerConfig{
		RabbitMQConsumerConfig: RabbitMQConsumerConfig{
			ExitOnError:      false,
			ConsumerID:       "",
			ConcurrencyLimit: 1,
			PrefetchCount:    4,
			AutoAck:          false,
			NoLocal:          false,
			NoWait:           true,
			BindingOptions: RabbitMQBindingOptions{
				RoutingKey: rabbitMQConfig.DeadLetterRoutingKey,
			},
			ExchangeOptions: RabbitMQExchangeOptions{
				Durable: true,
				Type:    enums.ExchangeTypeDirect,
				Name:    rabbitMQConfig.DeadLetterExchange,
			},
			QueueOptions: RabbitMQQueueOptions{
				Durable: true,
				Name:    rabbitMQConfig.DeadLetterExchange,
				Args:    nil,
			},
			Name: rabbitMQConfig.DeadLetterConsumer,
		},
		ConsumerMessageTypes: consumerMessageTypes,
		DeadLetterArgs:       NewDeadLetterArgs(rabbitMQConfig),
	}
}
