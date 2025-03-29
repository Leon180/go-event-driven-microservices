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
