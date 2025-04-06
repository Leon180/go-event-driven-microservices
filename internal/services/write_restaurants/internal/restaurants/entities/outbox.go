package entities

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type OutboxMessages []OutboxMessage

type OutboxMessage struct {
	ID         string             `gorm:"primaryKey"`
	MessageID  string             `gorm:"uniqueIndex"`
	Type       string             `gorm:"not null"`
	Payload    []byte             `gorm:"not null"` // marshalled json
	Status     enums.OutboxStatus `gorm:"not null"` // pending, published, failed
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Error      string
	RetryCount int
}

type UpdateOutboxMessage struct {
	ID         string
	Status     *enums.OutboxStatus
	Error      *string
	RetryCount *int
}

func (u *UpdateOutboxMessage) RemoveUnchangedFields(outboxMessage OutboxMessage) *UpdateOutboxMessage {
	if u == nil {
		return nil
	}
	if u.ID != outboxMessage.ID {
		return nil
	}
	if u.Status != nil && *u.Status == outboxMessage.Status {
		u.Status = nil
	}
	if u.Error != nil && *u.Error == outboxMessage.Error {
		u.Error = nil
	}
	if u.RetryCount != nil && *u.RetryCount == outboxMessage.RetryCount {
		u.RetryCount = nil
	}
	return u
}

func (u *UpdateOutboxMessage) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Status != nil {
		updateMap["status"] = *u.Status
	}
	if u.Error != nil {
		updateMap["error"] = *u.Error
	}
	if u.RetryCount != nil {
		updateMap["retry_count"] = *u.RetryCount
	}
	return updateMap
}
