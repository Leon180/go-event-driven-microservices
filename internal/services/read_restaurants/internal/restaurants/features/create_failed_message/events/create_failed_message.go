package events

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	serializers "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/services"
)

func NewCreateFailedMessageHandler(
	logger loggers.Logger,
	createFailedMessageService services.CreateFailedMessageHandler,
	messageSerializer serializers.MessageSerializer,
) *CreateFailedMessageHandle {
	return &CreateFailedMessageHandle{
		logger:                     logger,
		createFailedMessageService: createFailedMessageService,
		messageSerializer:          messageSerializer,
	}
}

type CreateFailedMessageHandle struct {
	logger                     loggers.Logger
	createFailedMessageService services.CreateFailedMessageHandler
	messageSerializer          serializers.MessageSerializer
}

func (h *CreateFailedMessageHandle) Handle(ctx context.Context, event types.MessageConsumeContext) error {
	serializationResult, err := h.messageSerializer.SerializeObject(event.Message())
	if err != nil {
		h.logger.Errorf("error in serializing message, error: {%v}", err)
		return err
	}
	aggregate := aggregates.FailedMessage{
		CorrelationID: event.CorrelationID(),
		MessageID:     event.MessageID(),
		Exchange:      event.Metadata().Get(string(enums.DeliveryHeaderExchange)).(string),
		RoutingKey:    event.Metadata().Get(string(enums.DeliveryHeaderRoutingKey)).(string),
		Queue:         event.Metadata().Get(string(enums.DeliveryHeaderQueue)).(string),
		MessageType:   event.Type(),
		ContentType:   event.ContentType().ToString(),
		Body:          serializationResult.Data,
		TimeStamp:     event.TimeStamp(),
		Error:         event.Metadata().Get(string(enums.DeliveryHeaderError)).(string),
		RetryCount:    int(event.Metadata().Get(string(enums.DeliveryHeaderRetryCount)).(int32)),
		CreatedAt:     time.Now(),
	}
	if err := h.createFailedMessageService.CreateFailedMessage(ctx, &aggregate); err != nil {
		h.logger.Errorf("error in sending CreateFailedMessage, error: {%v}", err)
		return err
	}
	h.logger.Info("CreateFailedMessage consumer handled.")
	return nil
}
