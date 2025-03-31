package rabbitmqconsumer

import (
	"fmt"
	"reflect"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/pipeline"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	customizereflect "github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type RabbitMQConsumerConfig struct {
	ExitOnError         bool
	ConsumerID          string
	Name                string
	ConsumerMessageType reflect.Type
	Pipelines           []pipeline.ConsumerPipeline
	Handlers            []consumer.ConsumerHandler
	ConcurrencyLimit    int
	PrefetchCount       int
	AutoAck             bool
	NoLocal             bool
	NoWait              bool
	BindingOptions      RabbitMQBindingOptions
	QueueOptions        RabbitMQQueueOptions
	ExchangeOptions     RabbitMQExchangeOptions
	MaxRetries          int
}

type RabbitMQBindingOptions struct {
	RoutingKey string
	Args       map[string]any
}

type RabbitMQExchangeOptions struct {
	Name       string
	Type       enums.ExchangeType
	AutoDelete bool
	Durable    bool
	Args       map[string]any
}

type RabbitMQQueueOptions struct {
	Name       string
	Durable    bool
	Exclusive  bool
	AutoDelete bool
	Args       map[string]any
}

func NewDefaultRabbitMQConsumerConfig(
	message types.Message,
) *RabbitMQConsumerConfig {
	name := fmt.Sprintf("%s_consumer", customizereflect.GetAnysTypeName(message))
	return &RabbitMQConsumerConfig{
		ExitOnError:      false,
		ConsumerID:       "",
		ConcurrencyLimit: 1,
		PrefetchCount:    4,
		AutoAck:          false,
		NoLocal:          false,
		NoWait:           true,
		BindingOptions: RabbitMQBindingOptions{
			RoutingKey: customizereflect.GetAnysTypeName(message),
		},
		ExchangeOptions: RabbitMQExchangeOptions{
			Durable: true,
			Type:    enums.ExchangeTypeTopic,
			Name:    customizereflect.GetAnysTypeName(message),
		},
		QueueOptions: RabbitMQQueueOptions{
			Durable: true,
			Name:    customizereflect.GetAnysTypeName(message),
		},
		ConsumerMessageType: customizereflect.GetAnysType(message),
		Name:                name,
		MaxRetries:          3,
	}
}
