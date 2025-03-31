package rabbitmqconsumer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/pipeline"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
)

// function to build rabbitmq consumer config
type DeadLetterRabbitMQConsumerConfigBuilderFunc func(builder DeadLetterRabbitMQConsumerConfigBuilder)

type DeadLetterRabbitMQConsumerConfigBuilder interface {
	SetHandlers(consumerHandler ...consumer.ConsumerHandler) DeadLetterRabbitMQConsumerConfigBuilder
	SetPipelines(pipeline ...pipeline.ConsumerPipeline) DeadLetterRabbitMQConsumerConfigBuilder
	SetExitOnError(exitOnError bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetName(name string) DeadLetterRabbitMQConsumerConfigBuilder
	SetAutoAck(ack bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetNoLocal(noLocal bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetNoWait(noWait bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetConcurrencyLimit(limit int) DeadLetterRabbitMQConsumerConfigBuilder
	SetPrefetchCount(count int) DeadLetterRabbitMQConsumerConfigBuilder
	SetConsumerId(consumerId string) DeadLetterRabbitMQConsumerConfigBuilder
	SetQueueName(queueName string) DeadLetterRabbitMQConsumerConfigBuilder
	SetDurable(durable bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetAutoDeleteQueue(autoDelete bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetExclusiveQueue(exclusive bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetQueueArgs(args map[string]any) DeadLetterRabbitMQConsumerConfigBuilder
	SetExchangeName(exchangeName string) DeadLetterRabbitMQConsumerConfigBuilder
	SetAutoDeleteExchange(autoDelete bool) DeadLetterRabbitMQConsumerConfigBuilder
	SetExchangeType(exchangeType enums.ExchangeType) DeadLetterRabbitMQConsumerConfigBuilder
	SetExchangeArgs(args map[string]any) DeadLetterRabbitMQConsumerConfigBuilder
	SetRoutingKey(routingKey string) DeadLetterRabbitMQConsumerConfigBuilder
	SetBindingArgs(args map[string]any) DeadLetterRabbitMQConsumerConfigBuilder
	Build() *DeadLetterRabbitMQConsumerConfig
}

func NewDeadLetterRabbitMQConsumerConfigBuilder(
	consumerMessageTypes []types.Message,
	rabbitMQConfig *rabbitmq.RabbitMQConfig,
) DeadLetterRabbitMQConsumerConfigBuilder {
	return &deadLetterRabbitMQConsumerConfigBuilder{
		rabbitmqConsumerConfig: NewDefaultDeadLetterRabbitMQConsumerConfig(consumerMessageTypes, rabbitMQConfig),
	}
}

type deadLetterRabbitMQConsumerConfigBuilder struct {
	rabbitmqConsumerConfig *DeadLetterRabbitMQConsumerConfig
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetHandlers(
	consumerHandler ...consumer.ConsumerHandler,
) DeadLetterRabbitMQConsumerConfigBuilder {
	if len(consumerHandler) > 0 {
		b.rabbitmqConsumerConfig.Handlers = append(b.rabbitmqConsumerConfig.Handlers, consumerHandler...)
	}
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetPipelines(
	pipeline ...pipeline.ConsumerPipeline,
) DeadLetterRabbitMQConsumerConfigBuilder {
	if len(pipeline) > 0 {
		b.rabbitmqConsumerConfig.Pipelines = append(b.rabbitmqConsumerConfig.Pipelines, pipeline...)
	}
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetExitOnError(
	exitOnError bool,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExitOnError = exitOnError
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetName(name string) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.Name = name
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetAutoAck(ack bool) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.AutoAck = ack
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetNoLocal(noLocal bool) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.NoLocal = noLocal
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetNoWait(noWait bool) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.NoWait = noWait
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetConcurrencyLimit(
	limit int,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ConcurrencyLimit = limit
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetPrefetchCount(count int) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.PrefetchCount = count
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetConsumerId(
	consumerID string,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ConsumerID = consumerID
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetQueueName(
	queueName string,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.Name = queueName
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetDurable(durable bool) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Durable = durable
	b.rabbitmqConsumerConfig.QueueOptions.Durable = durable
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetAutoDeleteQueue(
	autoDelete bool,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.AutoDelete = autoDelete
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetExclusiveQueue(
	exclusive bool,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.Exclusive = exclusive
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetQueueArgs(
	args map[string]any,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.Args = args
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetExchangeName(
	exchangeName string,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Name = exchangeName
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetAutoDeleteExchange(
	autoDelete bool,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.AutoDelete = autoDelete
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetExchangeType(
	exchangeType enums.ExchangeType,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Type = exchangeType
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetExchangeArgs(
	args map[string]any,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Args = args
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetRoutingKey(
	routingKey string,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.BindingOptions.RoutingKey = routingKey
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) SetBindingArgs(
	args map[string]any,
) DeadLetterRabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.BindingOptions.Args = args
	return b
}

func (b *deadLetterRabbitMQConsumerConfigBuilder) Build() *DeadLetterRabbitMQConsumerConfig {
	return b.rabbitmqConsumerConfig
}
