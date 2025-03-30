package rabbitmqconsumer

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/pipeline"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/connect"
	"github.com/iancoleman/strcase"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/ahmetb/go-linq/v3"
	"github.com/avast/retry-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConsumer(
	consumerConfig *RabbitMQConsumerConfig,
	rabbitmqConfig *rabbitmq.RabbitMQConfig,
	logger loggers.Logger,
	connection connect.AMQPConnection,
	messageSerializer serializers.MessageSerializer,
	consumedFuncs []consumer.ConsumedFunc,
) consumer.Consumer {
	return &rabbitMQConsumer{
		rabbitmqConsumerConfig: consumerConfig,
		rabbitmqConfig:         rabbitmqConfig,
		logger:                 logger,
		connection:             connection,
		messageSerializer:      messageSerializer,
		consumedFuncs:          consumedFuncs,
		deliveryRoutines:       make(chan struct{}, consumerConfig.ConcurrencyLimit),
		errorChan:              make(chan error),
	}
}

type rabbitMQConsumer struct {
	rabbitmqConsumerConfig *RabbitMQConsumerConfig
	rabbitmqConfig         *rabbitmq.RabbitMQConfig
	logger                 loggers.Logger
	connection             connect.AMQPConnection
	messageSerializer      serializers.MessageSerializer
	consumedFuncs          []consumer.ConsumedFunc
	channel                *amqp.Channel
	deliveryRoutines       chan struct{}
	errorChan              chan error
}

func (r *rabbitMQConsumer) Start(ctx context.Context) error {
	if r.connection == nil {
		return customizeerrors.RabbitmqConnectionError
	}

	exchangeName := r.rabbitmqConsumerConfig.ExchangeOptions.Name
	if r.rabbitmqConsumerConfig.ExchangeOptions.Name == "" {
		exchangeName = strcase.ToSnake(r.rabbitmqConsumerConfig.ConsumerMessageType.Name())
	}
	routingKey := r.rabbitmqConsumerConfig.BindingOptions.RoutingKey
	if r.rabbitmqConsumerConfig.BindingOptions.RoutingKey == "" {
		routingKey = strcase.ToSnake(r.rabbitmqConsumerConfig.ConsumerMessageType.Name())
	}
	queueName := r.rabbitmqConsumerConfig.QueueOptions.Name
	if r.rabbitmqConsumerConfig.QueueOptions.Name == "" {
		queueName = strcase.ToSnake(r.rabbitmqConsumerConfig.ConsumerMessageType.Name())
	}

	r.handleReconnectionEvent(ctx)

	// get a new channel on the connection - channel is unique for each consumer
	ch, err := r.connection.NewChannel()
	if err != nil {
		return customizeerrors.RabbitmqConnectionError
	}
	r.channel = ch

	// The prefetch count tells the Rabbit connection how many messages to retrieve from the server per request.
	prefetchCount := r.rabbitmqConsumerConfig.ConcurrencyLimit * r.rabbitmqConsumerConfig.PrefetchCount
	if err := r.channel.Qos(prefetchCount, 0, false); err != nil {
		return err
	}

	err = r.channel.ExchangeDeclare(
		exchangeName,
		r.rabbitmqConsumerConfig.ExchangeOptions.Type.ToString(),
		r.rabbitmqConsumerConfig.ExchangeOptions.Durable,
		r.rabbitmqConsumerConfig.ExchangeOptions.AutoDelete,
		false,
		r.rabbitmqConsumerConfig.NoWait,
		r.rabbitmqConsumerConfig.ExchangeOptions.Args)
	if err != nil {
		return err
	}

	_, err = r.channel.QueueDeclare(
		queueName,
		r.rabbitmqConsumerConfig.QueueOptions.Durable,
		r.rabbitmqConsumerConfig.QueueOptions.AutoDelete,
		r.rabbitmqConsumerConfig.QueueOptions.Exclusive,
		r.rabbitmqConsumerConfig.NoWait,
		r.rabbitmqConsumerConfig.QueueOptions.Args)
	if err != nil {
		return err
	}

	err = r.channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		r.rabbitmqConsumerConfig.NoWait,
		r.rabbitmqConsumerConfig.BindingOptions.Args)
	if err != nil {
		return err
	}

	delivering, err := r.channel.Consume(
		queueName,
		r.rabbitmqConsumerConfig.ConsumerID,
		r.rabbitmqConsumerConfig.AutoAck, // When autoAck (also known as noAck) is true, the server will acknowledge deliveries to this consumer prior to writing the delivery to the network. When autoAck is true, the consumer should not call Delivery.Ack.
		r.rabbitmqConsumerConfig.QueueOptions.Exclusive,
		r.rabbitmqConsumerConfig.NoLocal,
		r.rabbitmqConsumerConfig.NoWait,
		nil,
	)
	if err != nil {
		return err
	}

	chClosedCh := make(chan *amqp.Error, 1)
	r.channel.NotifyClose(chClosedCh)

	for i := range r.rabbitmqConsumerConfig.ConcurrencyLimit {
		r.logger.Infof("Processing messages on thread %d", i)
		go func() {
			for {
				select {
				case <-ctx.Done():
					r.logger.Info("shutting down consumer")
					return
				case amqErr := <-chClosedCh:
					r.logger.Errorf("AMQP Channel closed due to: %s", amqErr)
					chClosedCh = make(chan *amqp.Error, 1)
					r.channel.NotifyClose(chClosedCh)
				case delivery, ok := <-delivering:
					if !ok {
						r.logger.Info("consumer connection dropped")
						return
					}
					r.handleReceived(ctx, delivery)
				}
			}
		}()
	}

	return nil
}

func (r *rabbitMQConsumer) Stop() error {
	defer func() {
		if r.channel == nil || r.channel.IsClosed() {
			return
		}
		r.channel.Cancel(r.rabbitmqConsumerConfig.ConsumerID, false)
		r.channel.Close()
	}()

	done := make(chan struct{}, 1)
	go func() {
		for {
			if len(r.deliveryRoutines) != 0 {
				continue
			}
			done <- struct{}{}
		}
	}()
	<-done

	return nil
}

func (r *rabbitMQConsumer) BindHandler(handler ...consumer.ConsumerHandler) {
	r.rabbitmqConsumerConfig.Handlers = append(r.rabbitmqConsumerConfig.Handlers, handler...)
}

func (r *rabbitMQConsumer) Name() string {
	return r.rabbitmqConsumerConfig.Name
}

func (r *rabbitMQConsumer) handleReconnectionEvent(ctx context.Context) {
	go func() {
		for range r.connection.ReconnectedEvent() {
			r.logger.Info("restarting consumer")
			err := r.Start(ctx)
			if err != nil {
				r.logger.Error("restarting consumer failed with error: %v", err)
			}
			r.logger.Info("restarting consumer finished successfully")
		}
	}()
}

func (r *rabbitMQConsumer) handleReceived(ctx context.Context, delivery amqp.Delivery) {
	// for ensuring our handlers execute completely after shutdown
	r.deliveryRoutines <- struct{}{}
	defer func() { <-r.deliveryRoutines }()

	consumeContext, err := r.createMessageConsumeContext(delivery)
	if err != nil {
		return
	}

	if r.rabbitmqConsumerConfig.AutoAck {
		r.handle(ctx, nil, nil, consumeContext)
		return
	}

	customizeAck := func() {
		if err := delivery.Ack(false); err != nil {
			r.logger.Error("error sending ACK to RabbitMQ consumer: %v", err)
			return
		}
		for _, cousumedFunc := range r.consumedFuncs {
			if cousumedFunc == nil {
				continue
			}
			cousumedFunc(consumeContext.Message())
		}
	}
	customizeNack := func() {
		if err := delivery.Nack(false, true); err != nil {
			r.logger.Error("error in sending Nack to RabbitMQ consumer: %v", err)
			return
		}
	}
	r.handle(ctx, &customizeAck, &customizeNack, consumeContext)
}

func (r *rabbitMQConsumer) handle(
	ctx context.Context,
	customizeAck *func(),
	customizeNack *func(),
	messageConsumeContext types.MessageConsumeContext,
) {
	for _, handler := range r.rabbitmqConsumerConfig.Handlers {
		err := r.handleWithRetry(ctx, handler, messageConsumeContext)
		if err != nil {
			r.logger.Error("error handling consume message, prepare for nacking message")
			if customizeNack != nil {
				(*customizeNack)()
			}
			return
		}
	}

	if customizeAck != nil {
		(*customizeAck)()
	}
}

func (r *rabbitMQConsumer) handleWithRetry(
	ctx context.Context,
	handler consumer.ConsumerHandler,
	messageConsumeContext types.MessageConsumeContext,
) error {
	r.logger.Info("handling message, correlation id: %s, message id: %s, message type: %s",
		messageConsumeContext.CorrelationID(),
		messageConsumeContext.MessageID(),
		messageConsumeContext.Type(),
	)
	return retry.Do(func() error {
		handleFunc := func(ctx context.Context) error {
			return handler.Handle(ctx, messageConsumeContext)
		}
		if len(r.rabbitmqConsumerConfig.Pipelines) > 0 {
			reversePipelines := pipeline.ReversePipelinesOrder(r.rabbitmqConsumerConfig.Pipelines...)
			handleFunc = linq.
				From(reversePipelines).
				AggregateWithSeedT(handleFunc, func(next func(ctx context.Context) error, pipe pipeline.ConsumerPipeline) func(ctx context.Context) error {
					return func(ctx context.Context) error {
						return pipe.Handle(ctx, messageConsumeContext, next)
					}
				}).(func(ctx context.Context) error)
		}
		return handleFunc(ctx)
	},
		retry.Attempts(uint(r.rabbitmqConfig.RetryAttempts)),
		retry.Delay(time.Duration(r.rabbitmqConfig.RetryDelay)*time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.Context(ctx),
	)
}

func (r *rabbitMQConsumer) createMessageConsumeContext(
	delivery amqp.Delivery,
) (types.MessageConsumeContext, error) {
	message := r.deserializeData(
		delivery.Body,
		delivery.Type,
		enums.ContentType(delivery.ContentType),
	)

	var meta metadatas.Metadata
	if delivery.Headers != nil {
		meta = metadatas.MapToMetadata(delivery.Headers)
	}

	consumeContext := types.NewMessageConsumeContext(
		message,
		delivery.MessageId,
		meta,
		enums.ContentType(delivery.ContentType),
		delivery.Type,
		delivery.DeliveryTag,
		delivery.CorrelationId,
		delivery.Timestamp,
	)
	return consumeContext, nil
}

func (r *rabbitMQConsumer) deserializeData(body []byte, eventType string, contentType enums.ContentType) types.Message {
	if contentType == "" {
		contentType = enums.ContentTypeJSON
	}
	if len(body) == 0 {
		return nil
	}
	switch contentType {
	case enums.ContentTypeJSON:
		deserialize, err := r.messageSerializer.Deserialize(
			body,
			eventType,
			contentType,
		) // or this to explicit type deserialization
		if err != nil {
			r.logger.Errorf("error in deserilizng of type '%s' in the consumer: %v", eventType, err)
			return nil
		}
		return deserialize
	default:
		return nil
	}
}
