package aggregates

import (
	"time"
)

type FailedMessages []FailedMessage

type FailedMessage struct {
	CorrelationID string
	MessageID     string
	Exchange      string
	RoutingKey    string
	Queue         string
	MessageType   string
	ContentType   string
	Body          []byte
	TimeStamp     time.Time
	Error         string
	RetryCount    int
	CreatedAt     time.Time
}
