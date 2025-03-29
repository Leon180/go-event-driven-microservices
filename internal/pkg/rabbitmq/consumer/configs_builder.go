package rabbitmqconsumer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	consumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/pipeline"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

// function to build rabbitmq consumer config
type RabbitMQConsumerConfigBuilderFunc func(builder RabbitMQConsumerConfigBuilder)

type RabbitMQConsumerConfigBuilder interface {
	SetHandlers(consumerHandler ...consumer.ConsumerHandler) RabbitMQConsumerConfigBuilder
	SetPipelines(pipeline ...pipeline.ConsumerPipeline) RabbitMQConsumerConfigBuilder
	SetExitOnError(exitOnError bool) RabbitMQConsumerConfigBuilder
	SetName(name string) RabbitMQConsumerConfigBuilder
	SetAutoAck(ack bool) RabbitMQConsumerConfigBuilder
	SetNoLocal(noLocal bool) RabbitMQConsumerConfigBuilder
	SetNoWait(noWait bool) RabbitMQConsumerConfigBuilder
	SetConcurrencyLimit(limit int) RabbitMQConsumerConfigBuilder
	SetPrefetchCount(count int) RabbitMQConsumerConfigBuilder
	SetConsumerId(consumerId string) RabbitMQConsumerConfigBuilder
	SetQueueName(queueName string) RabbitMQConsumerConfigBuilder
	SetDurable(durable bool) RabbitMQConsumerConfigBuilder
	SetAutoDeleteQueue(autoDelete bool) RabbitMQConsumerConfigBuilder
	SetExclusiveQueue(exclusive bool) RabbitMQConsumerConfigBuilder
	SetQueueArgs(args map[string]any) RabbitMQConsumerConfigBuilder
	SetExchangeName(exchangeName string) RabbitMQConsumerConfigBuilder
	SetAutoDeleteExchange(autoDelete bool) RabbitMQConsumerConfigBuilder
	SetExchangeType(exchangeType enums.ExchangeType) RabbitMQConsumerConfigBuilder
	SetExchangeArgs(args map[string]any) RabbitMQConsumerConfigBuilder
	SetRoutingKey(routingKey string) RabbitMQConsumerConfigBuilder
	SetBindingArgs(args map[string]any) RabbitMQConsumerConfigBuilder
	Build() *RabbitMQConsumerConfig
}

func NewRabbitMQConsumerConfigBuilder(message types.Message) RabbitMQConsumerConfigBuilder {
	return &rabbitMQConsumerConfigBuilder{
		rabbitmqConsumerConfig: NewDefaultRabbitMQConsumerConfig(message),
	}
}

type rabbitMQConsumerConfigBuilder struct {
	rabbitmqConsumerConfig *RabbitMQConsumerConfig
}

func (b *rabbitMQConsumerConfigBuilder) SetHandlers(
	consumerHandler ...consumer.ConsumerHandler,
) RabbitMQConsumerConfigBuilder {
	if len(consumerHandler) > 0 {
		b.rabbitmqConsumerConfig.Handlers = append(b.rabbitmqConsumerConfig.Handlers, consumerHandler...)
	}
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetPipelines(
	pipeline ...pipeline.ConsumerPipeline,
) RabbitMQConsumerConfigBuilder {
	if len(pipeline) > 0 {
		b.rabbitmqConsumerConfig.Pipelines = append(b.rabbitmqConsumerConfig.Pipelines, pipeline...)
	}
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetExitOnError(exitOnError bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExitOnError = exitOnError
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetName(name string) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.Name = name
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetAutoAck(ack bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.AutoAck = ack
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetNoLocal(noLocal bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.NoLocal = noLocal
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetNoWait(noWait bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.NoWait = noWait
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetConcurrencyLimit(limit int) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ConcurrencyLimit = limit
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetPrefetchCount(count int) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.PrefetchCount = count
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetConsumerId(consumerID string) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ConsumerID = consumerID
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetQueueName(queueName string) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.Name = queueName
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetDurable(durable bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Durable = durable
	b.rabbitmqConsumerConfig.QueueOptions.Durable = durable
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetAutoDeleteQueue(autoDelete bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.AutoDelete = autoDelete
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetExclusiveQueue(exclusive bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.Exclusive = exclusive
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetQueueArgs(args map[string]any) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.QueueOptions.Args = args
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetExchangeName(exchangeName string) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Name = exchangeName
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetAutoDeleteExchange(autoDelete bool) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.AutoDelete = autoDelete
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetExchangeType(exchangeType enums.ExchangeType) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Type = exchangeType
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetExchangeArgs(args map[string]any) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.ExchangeOptions.Args = args
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetRoutingKey(routingKey string) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.BindingOptions.RoutingKey = routingKey
	return b
}

func (b *rabbitMQConsumerConfigBuilder) SetBindingArgs(args map[string]any) RabbitMQConsumerConfigBuilder {
	b.rabbitmqConsumerConfig.BindingOptions.Args = args
	return b
}

func (b *rabbitMQConsumerConfigBuilder) Build() *RabbitMQConsumerConfig {
	return b.rabbitmqConsumerConfig
}
