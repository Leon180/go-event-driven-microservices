package consumer

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
)

type ConsumerHandler interface {
	Handle(ctx context.Context, consumeContext types.MessageConsumeContext) error
}
