package documents

import (
	"time"
)

type FailedMessages []FailedMessage

type FailedMessage struct {
	CorrelationID string    `bson:"correlation_id"`
	MessageID     string    `bson:"message_id,unique"`
	Exchange      string    `bson:"exchange"`
	RoutingKey    string    `bson:"routing_key"`
	Queue         string    `bson:"queue"`
	MessageType   string    `bson:"message_type"`
	ContentType   string    `bson:"content_type"`
	Body          []byte    `bson:"body"`
	DeliveryMode  int       `bson:"delivery_mode"`
	TimeStamp     time.Time `bson:"time_stamp"`
	Error         string    `bson:"error"`
	RetryCount    int       `bson:"retry_count"`
	CreatedAt     time.Time `bson:"created_at"`
}
