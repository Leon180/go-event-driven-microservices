package producer

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
)

type Producer interface {
	PublishMessage(
		ctx context.Context,
		message types.Message,
		meta metadatas.Metadata,
		topicOrExchangeName *string,
	) error
	Produced(producedFuncs ...ProducedFunc)
}

type ProducedFunc func(message types.Message)
