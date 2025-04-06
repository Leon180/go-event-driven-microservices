package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
)

type UpdateOutboxMessages interface {
	CreateOutboxMessages(ctx context.Context, entities entities.OutboxMessages) error
	UpdateOutboxMessage(ctx context.Context, update *entities.UpdateOutboxMessage) error
	DeleteOutboxMessages(ctx context.Context, ids []string) error
}

type UpdateOutboxMessagesWithTransaction interface {
	customizegorm.Transaction
	UpdateOutboxMessages
}

type ListOutboxMessages interface {
	ListOutboxMessages(ctx context.Context, status []enums.OutboxStatus, retry *int) (entities.OutboxMessages, error)
}
