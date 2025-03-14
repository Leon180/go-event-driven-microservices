package messageheader

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
)

const (
	MessageID     string = "message-id"
	CorrelationID string = "correlation-id"
	Name          string = "name"
	MessageType   string = "message-type"
	ContentType   string = "content-type"
	TimeStamp     string = "time-stamp"
)

func GetCorrelationID(m metadatas.Metadata) string {
	return m.Get(CorrelationID).(string)
}

func SetCorrelationID(m metadatas.Metadata, val string) {
	m.Set(CorrelationID, val)
}

func GetMessageID(m metadatas.Metadata) string {
	return m.Get(MessageID).(string)
}

func SetMessageID(m metadatas.Metadata, val string) {
	m.Set(MessageID, val)
}

func GetMessageName(m metadatas.Metadata) string {
	return m.Get(Name).(string)
}

func SetMessageName(m metadatas.Metadata, val string) {
	m.Set(Name, val)
}

func GetMessageType(m metadatas.Metadata) string {
	return m.Get(MessageType).(string)
}

func SetMessageType(m metadatas.Metadata, val string) {
	m.Set(MessageType, val)
}

func SetContentType(m metadatas.Metadata, val string) {
	m.Set(ContentType, val)
}

func GetContentType(m metadatas.Metadata) string {
	return m.Get(ContentType).(string)
}

func GetTimeStamp(m metadatas.Metadata) time.Time {
	return m.Get(TimeStamp).(time.Time)
}

func SetTimeStamp(m metadatas.Metadata, val time.Time) {
	m.Set(TimeStamp, val)
}
