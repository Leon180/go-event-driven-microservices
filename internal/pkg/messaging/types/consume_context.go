package types

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
)

type MessageConsumeContext interface {
	MessageID() string
	Type() string
	TimeStamp() time.Time
	CorrelationID() string
	ContentType() enums.ContentType
	DeliveryTag() uint64

	Metadata() metadatas.Metadata
	Message() Message
}

func NewMessageConsumeContext(
	message Message,
	messageID string,
	metadata metadatas.Metadata,
	contentType enums.ContentType,
	messageType string,
	tag uint64,
	correlationID string,
	timeStamp time.Time,
) MessageConsumeContext {
	return &messageConsumeContextImpl{
		message:       message,
		messageID:     messageID,
		metadata:      metadata,
		contentType:   contentType,
		messageType:   messageType,
		tag:           tag,
		correlationID: correlationID,
		timeStamp:     timeStamp,
	}
}

type messageConsumeContextImpl struct {
	messageID     string
	messageType   string
	timeStamp     time.Time
	contentType   enums.ContentType
	tag           uint64
	correlationID string

	metadata metadatas.Metadata
	message  Message
}

func (m *messageConsumeContextImpl) MessageID() string {
	return m.messageID
}

func (m *messageConsumeContextImpl) Type() string {
	return m.messageType
}

func (m *messageConsumeContextImpl) TimeStamp() time.Time {
	return m.timeStamp
}

func (m *messageConsumeContextImpl) CorrelationID() string {
	return m.correlationID
}

func (m *messageConsumeContextImpl) ContentType() enums.ContentType {
	return m.contentType
}

func (m *messageConsumeContextImpl) DeliveryTag() uint64 {
	return m.tag
}

func (m *messageConsumeContextImpl) Metadata() metadatas.Metadata {
	return m.metadata
}

func (m *messageConsumeContextImpl) Message() Message {
	return m.message
}
