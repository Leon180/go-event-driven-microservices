package enums

import amqp "github.com/rabbitmq/amqp091-go"

type ExchangeType string

const (
	ExchangeTypeDirect ExchangeType = amqp.ExchangeDirect
	ExchangeTypeFanout ExchangeType = amqp.ExchangeFanout
	ExchangeTypeTopic  ExchangeType = amqp.ExchangeTopic
)

func (e ExchangeType) ToString() string {
	return string(e)
}

type DeliveryHeader string

const (
	DeliveryHeaderDeadLetterExchange   DeliveryHeader = "x-dead-letter-exchange"
	DeliveryHeaderDeadLetterRoutingKey DeliveryHeader = "x-dead-letter-routing-key"
	DeliveryHeaderRetryCount           DeliveryHeader = "x-retry-count"
	DeliveryHeaderError                DeliveryHeader = "x-error"
	DeliveryHeaderExchange             DeliveryHeader = "x-exchange"
	DeliveryHeaderRoutingKey           DeliveryHeader = "x-routing-key"
	DeliveryHeaderQueue                DeliveryHeader = "x-queue"
)

func (d DeliveryHeader) ToString() string {
	return string(d)
}
