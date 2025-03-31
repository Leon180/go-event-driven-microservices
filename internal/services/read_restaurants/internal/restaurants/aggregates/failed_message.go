package aggregates

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/samber/lo"
)

type FailedMessages []FailedMessage

func (f *FailedMessages) ToDocuments() documents.FailedMessages {
	return lo.Map(*f, func(failedMessage FailedMessage, _ int) documents.FailedMessage {
		return *failedMessage.ToDocument()
	})
}

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

func (f *FailedMessage) ToDocument() *documents.FailedMessage {
	return &documents.FailedMessage{
		CorrelationID: f.CorrelationID,
		MessageID:     f.MessageID,
		Exchange:      f.Exchange,
		RoutingKey:    f.RoutingKey,
		Queue:         f.Queue,
		MessageType:   f.MessageType,
		ContentType:   f.ContentType,
		Body:          f.Body,
		TimeStamp:     f.TimeStamp,
		Error:         f.Error,
		RetryCount:    f.RetryCount,
		CreatedAt:     f.CreatedAt,
	}
}

type FailedMessageDocuments documents.FailedMessages

func (f FailedMessageDocuments) ToAggregate() FailedMessages {
	return lo.Map(f, func(failedMessageDocument documents.FailedMessage, _ int) FailedMessage {
		d := FailedMessageDocument(failedMessageDocument)
		return *d.ToAggregate()
	})
}

type FailedMessageDocument documents.FailedMessage

func (f *FailedMessageDocument) ToAggregate() *FailedMessage {
	return &FailedMessage{
		CorrelationID: f.CorrelationID,
		MessageID:     f.MessageID,
		Exchange:      f.Exchange,
		RoutingKey:    f.RoutingKey,
		Queue:         f.Queue,
		MessageType:   f.MessageType,
		ContentType:   f.ContentType,
		Body:          f.Body,
		TimeStamp:     f.TimeStamp,
		Error:         f.Error,
		RetryCount:    f.RetryCount,
		CreatedAt:     f.CreatedAt,
	}
}
