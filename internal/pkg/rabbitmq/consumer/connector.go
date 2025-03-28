package rabbitmqconsumer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

type RabbitMQConsumerConnector interface {
	consumer.ConsumerConnector
	// ConnectRabbitMQConsumer Add a new consumer to existing message type consumers. if there is no consumer, will create a new consumer for the message type
	ConnectRabbitMQConsumer(
		message types.Message,
		consumerBuilderFunc RabbitMQConsumerConfigBuilderFunc,
	)
}
