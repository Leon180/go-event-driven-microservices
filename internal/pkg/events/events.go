package events

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

func NewEvent(id string, eventType string) Event {
	return &EventImpl{
		EventID:    id,
		OccurredAt: time.Now(),
		EventType:  eventType,
	}
}

type Event interface {
	GetEventID() string
	GetOccurredAt() time.Time
	GetEventTypeName() string
	GetEventFullTypeName() string
}

func IsEvent(obj any) bool {
	if _, ok := obj.(Event); ok {
		return true
	}
	return false
}

type EventImpl struct {
	EventID    string    `json:"event_id"`
	EventType  string    `json:"event_type"`
	OccurredAt time.Time `json:"occurred_at"`
}

func (e *EventImpl) GetEventID() string {
	return e.EventID
}

func (e *EventImpl) GetEventType() string {
	return e.EventType
}

func (e *EventImpl) GetOccurredAt() time.Time {
	return e.OccurredAt
}

func (e *EventImpl) GetEventTypeName() string {
	return reflect.GetAnysTypeName(e)
}

func (e *EventImpl) GetEventFullTypeName() string {
	return reflect.GetAnysFullTypeName(e)
}
