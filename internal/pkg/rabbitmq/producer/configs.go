package producer

import (
	"reflect"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	customizereflect "github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type RabbitMQProducerConfig struct {
	ProducerMessageType reflect.Type
	DeliveryMode        uint8
	Priority            uint8
	AppId               string
	Expiration          string
	ReplyTo             string
	ContentEncoding     string
	BindingOptions      RabbitMQBindingOptions
	ExchangeOptions     RabbitMQExchangeOptions
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

func NewDefaultRabbitMQProducerConfig(
	message types.Message,
) *RabbitMQProducerConfig {
	return &RabbitMQProducerConfig{
		ExchangeOptions: RabbitMQExchangeOptions{
			Durable: true,
			Type:    enums.ExchangeTypeTopic,
			Name:    customizereflect.GetAnysTypeName(message),
		},
		DeliveryMode: 2,
		BindingOptions: RabbitMQBindingOptions{
			RoutingKey: customizereflect.GetAnysTypeName(message),
			Args:       nil,
		},
		ProducerMessageType: customizereflect.GetAnysType(message),
	}
}
