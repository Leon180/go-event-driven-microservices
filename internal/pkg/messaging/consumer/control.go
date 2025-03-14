package consumer

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

type ConsumedFunc func(message types.Message)

type ConsumerControl interface {
	Start(ctx context.Context) error // Start starts all consumers
	Stop() error                     // Stop stops all consumers
	Consumed(consumedFuncs ...ConsumedFunc)
}
