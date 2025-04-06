package outbox

import (
	"context"
	"log"
	"time"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

func NewOutboxOperator(
	producer producer.Producer,
	repository repositories.ListOutboxMessages,
	updateOutboxMessageRepository repositories.UpdateOutboxMessages,
	config *OutboxConfig,
	logger loggers.Logger,
	serializer serializers.MessageSerializer,
) *OutboxOperator {
	return &OutboxOperator{
		producer:                      producer,
		repository:                    repository,
		updateOutboxMessageRepository: updateOutboxMessageRepository,
		config:                        config,
		logger:                        logger,
		serializer:                    serializer,
	}
}

type OutboxOperator struct {
	producer                      producer.Producer
	repository                    repositories.ListOutboxMessages
	updateOutboxMessageRepository repositories.UpdateOutboxMessages
	config                        *OutboxConfig
	logger                        loggers.Logger
	serializer                    serializers.MessageSerializer
}

func (p *OutboxOperator) Start(ctx context.Context) error {
	ticker := time.NewTicker(p.config.Interval * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := p.processPendingMessages(ctx); err != nil {
				log.Printf("Error processing pending messages: %v", err)
			}
		}
	}
}

func (p *OutboxOperator) processPendingMessages(ctx context.Context) error {
	// Find pending messages with retry count less than max
	outboxMessages, err := p.repository.ListOutboxMessages(
		ctx,
		[]enums.OutboxStatus{enums.OutboxStatusPending, enums.OutboxStatusFailed},
		lo.ToPtr(p.config.MaxRetries),
	)
	if err != nil {
		return err
	}

	for _, outboxMessage := range outboxMessages {
		if err := p.publishMessage(ctx, outboxMessage); err != nil {
			p.logger.Error("Error publishing message %s: %v", outboxMessage.MessageID, err)
			continue
		}
	}

	return nil
}

func (p *OutboxOperator) publishMessage(ctx context.Context, outboxMessage entities.OutboxMessage) error {
	// Publish to RabbitMQ
	message, err := p.serializer.Deserialize(outboxMessage.Payload, outboxMessage.Type, enums.ContentTypeJSON)
	if err != nil {
		return err
	}

	if err = p.producer.PublishMessage(ctx, message, nil, nil); err != nil {
		// Update message status to failed
		updateOutboxMessage := &entities.UpdateOutboxMessage{
			ID:         outboxMessage.ID,
			Status:     lo.ToPtr(enums.OutboxStatusFailed),
			Error:      lo.ToPtr(err.Error()),
			RetryCount: lo.ToPtr(outboxMessage.RetryCount + 1),
		}
		if err := p.updateOutboxMessageRepository.UpdateOutboxMessage(ctx, updateOutboxMessage); err != nil {
			p.logger.Error("Error updating message status: %v", err)
			return err
		}
		return err
	}

	// Update message status to published
	updateOutboxMessage := &entities.UpdateOutboxMessage{
		ID:     outboxMessage.ID,
		Status: lo.ToPtr(enums.OutboxStatusPublished),
	}
	if err := p.updateOutboxMessageRepository.UpdateOutboxMessage(ctx, updateOutboxMessage); err != nil {
		p.logger.Error("Error updating message status: %v", err)
		return err
	}

	return nil
}
