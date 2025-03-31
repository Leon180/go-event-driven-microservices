package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
)

type SearchFailedMessagesMongo interface {
	SearchFailedMessages(
		ctx context.Context,
		searchFailedMessages *dtos.SearchFailedMessages,
	) (aggregates.FailedMessages, error)
}

type ReadFailedMessageMongo interface {
	ReadFailedMessage(ctx context.Context, messageID string) (*aggregates.FailedMessage, error)
}

type CreateFailedMessageMongo interface {
	Create(ctx context.Context, failedMessage *aggregates.FailedMessage) error
}
