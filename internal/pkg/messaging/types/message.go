package types

import (
	"time"
)

type Message interface {
	ID() string
	TimeStamp() time.Time
	Type() string
}

func NewMessage(messageID string, messageType string) Message {
	return &MessageImpl{
		MessageID:        messageID,
		MessageTimeStamp: time.Now(),
		MessageType:      messageType,
	}
}

type MessageImpl struct {
	MessageID        string    `json:"message_id,omitempty"`
	MessageTimeStamp time.Time `json:"message_time_stamp"`
	MessageType      string    `json:"message_type"`
}

func (m *MessageImpl) ID() string {
	return m.MessageID
}

func (m *MessageImpl) TimeStamp() time.Time {
	return m.MessageTimeStamp
}

func (m *MessageImpl) Type() string {
	return m.MessageType
}
