package consumer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

type ConsumerConnector interface {
	// ConnectConsumerHandler Add handler to existing consumer. creates new consumer if not exist
	ConnectConsumerHandler(message types.Message, consumerHandler ConsumerHandler)
	// ConnectConsumer Add a new consumer to existing message type consumers. if there is no consumer, will create a new consumer for the message
	ConnectConsumer(message types.Message, consumer Consumer)
}
