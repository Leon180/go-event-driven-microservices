package pipeline

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/samber/lo/mutable"
)

type ConsumerPipeline interface {
	Handle(ctx context.Context, consumerContext types.MessageConsumeContext, next func(ctx context.Context) error) error
}

func ReversePipelinesOrder(pipes ...ConsumerPipeline) []ConsumerPipeline {
	reversedPipelines := pipes
	mutable.Reverse(reversedPipelines)
	return reversedPipelines
}
