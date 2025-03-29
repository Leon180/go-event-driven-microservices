package producer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

type RabbitMQProducerConfigBuilderFunc func(builder RabbitMQProducerConfigBuilder)

type RabbitMQProducerConfigBuilder interface {
	SetDurable(durable bool) RabbitMQProducerConfigBuilder
	SetAutoDeleteExchange(autoDelete bool) RabbitMQProducerConfigBuilder
	SetExchangeType(exchangeType enums.ExchangeType) RabbitMQProducerConfigBuilder
	SetExchangeName(exchangeName string) RabbitMQProducerConfigBuilder
	SetRoutingKey(routingKey string) RabbitMQProducerConfigBuilder
	SetExchangeArgs(args map[string]any) RabbitMQProducerConfigBuilder
	SetDeliveryMode(deliveryMode uint8) RabbitMQProducerConfigBuilder
	SetPriority(priority uint8) RabbitMQProducerConfigBuilder
	SetAppId(appId string) RabbitMQProducerConfigBuilder
	SetExpiration(expiration string) RabbitMQProducerConfigBuilder
	SetReplyTo(replyTo string) RabbitMQProducerConfigBuilder
	SetContentEncoding(contentEncoding string) RabbitMQProducerConfigBuilder
	Build() *RabbitMQProducerConfig
}

type rabbitMQProducerConfigBuilder struct {
	rabbitmqProducerConfig *RabbitMQProducerConfig
}

func NewRabbitMQProducerConfigBuilder(
	messageType types.Message,
) RabbitMQProducerConfigBuilder {
	return &rabbitMQProducerConfigBuilder{
		rabbitmqProducerConfig: NewDefaultRabbitMQProducerConfig(messageType),
	}
}

func (b *rabbitMQProducerConfigBuilder) SetDurable(durable bool) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ExchangeOptions.Durable = durable
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetAutoDeleteExchange(autoDelete bool) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ExchangeOptions.AutoDelete = autoDelete
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetExchangeType(exchangeType enums.ExchangeType) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ExchangeOptions.Type = exchangeType
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetRoutingKey(routingKey string) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.BindingOptions.RoutingKey = routingKey
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetExchangeName(exchangeName string) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ExchangeOptions.Name = exchangeName
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetExchangeArgs(args map[string]any) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ExchangeOptions.Args = args
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetDeliveryMode(deliveryMode uint8) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.DeliveryMode = deliveryMode
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetPriority(priority uint8) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.Priority = priority
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetAppId(appId string) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.AppId = appId
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetExpiration(expiration string) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.Expiration = expiration
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetReplyTo(replyTo string) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ReplyTo = replyTo
	return b
}

func (b *rabbitMQProducerConfigBuilder) SetContentEncoding(contentEncoding string) RabbitMQProducerConfigBuilder {
	b.rabbitmqProducerConfig.ContentEncoding = contentEncoding
	return b
}

func (b *rabbitMQProducerConfigBuilder) Build() *RabbitMQProducerConfig {
	return b.rabbitmqProducerConfig
}
