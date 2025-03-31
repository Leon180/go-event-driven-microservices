package rabbitmqconsumer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/connect"
)

type ConsumerFactory interface {
	CreateConsumer(config *RabbitMQConsumerConfig, consumedFuncs []consumer.ConsumedFunc) consumer.Consumer
	CreateDeadLetterConsumer(config *RabbitMQConsumerConfig, consumedFuncs []consumer.ConsumedFunc) consumer.Consumer
	Connection() connect.AMQPConnection
}

func NewConsumerFactory(
	rabbitmqConfig *rabbitmq.RabbitMQConfig,
	connection connect.AMQPConnection,
	eventSerializer serializers.MessageSerializer,
	logger loggers.Logger,
) ConsumerFactory {
	return &consumerFactory{
		rabbitmqConfig:  rabbitmqConfig,
		connection:      connection,
		eventSerializer: eventSerializer,
		logger:          logger,
	}
}

type consumerFactory struct {
	rabbitmqConfig  *rabbitmq.RabbitMQConfig
	connection      connect.AMQPConnection
	eventSerializer serializers.MessageSerializer
	logger          loggers.Logger
}

func (c *consumerFactory) CreateConsumer(
	consumerConfig *RabbitMQConsumerConfig,
	consumedFuncs []consumer.ConsumedFunc,
) consumer.Consumer {
	return NewRabbitMQConsumer(
		consumerConfig,
		c.rabbitmqConfig,
		c.logger,
		c.connection,
		c.eventSerializer,
		consumedFuncs,
	)
}

func (c *consumerFactory) Connection() connect.AMQPConnection {
	return c.connection
}

func (c *consumerFactory) CreateDeadLetterConsumer(
	consumerConfig *RabbitMQConsumerConfig,
	consumedFuncs []consumer.ConsumedFunc,
) consumer.Consumer {
	return NewDeadLetterRabbitMQConsumer(
		consumerConfig,
		c.rabbitmqConfig,
		c.logger,
		c.connection,
		c.eventSerializer,
		consumedFuncs,
	)
}
