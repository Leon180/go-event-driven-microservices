package bus

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
)

type Bus interface {
	producer.Producer
	consumer.ConsumerConnector
	consumer.ConsumerControl
}
