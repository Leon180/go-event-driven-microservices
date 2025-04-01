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
	exchangeName           string
	queueName              string
	routingKey             string
	chClosedCh             chan *amqp.Error
}

func (r *rabbitMQConsumer) Start(ctx context.Context) error {
	var err error
	if r.connection == nil {
		return customizeerrors.RabbitmqConnectionError
	}
	r.exchangeName = r.rabbitmqConsumerConfig.ExchangeOptions.Name
	if r.rabbitmqConsumerConfig.ExchangeOptions.Name == "" {
		r.exchangeName = strcase.ToSnake(r.rabbitmqConsumerConfig.ConsumerMessageType.Name())
	}
	r.routingKey = r.rabbitmqConsumerConfig.BindingOptions.RoutingKey
	if r.rabbitmqConsumerConfig.BindingOptions.RoutingKey == "" {
		r.routingKey = strcase.ToSnake(r.rabbitmqConsumerConfig.ConsumerMessageType.Name())
	}
	r.queueName = r.rabbitmqConsumerConfig.QueueOptions.Name
	if r.rabbitmqConsumerConfig.QueueOptions.Name == "" {
		r.queueName = strcase.ToSnake(r.rabbitmqConsumerConfig.ConsumerMessageType.Name())
	}

	// get a new channel on the connection - channel is unique for each consumer
	if r.channel, err = r.connection.NewChannel(); err != nil {
		r.logger.Errorf("error creating new channel, error: %v", err)
		return customizeerrors.RabbitmqConnectionError
	}

	// The prefetch count tells the Rabbit connection how many messages to retrieve from the server per request.
	if err = r.channel.Qos(r.rabbitmqConsumerConfig.ConcurrencyLimit*r.rabbitmqConsumerConfig.PrefetchCount, 0, false); err != nil {
		r.logger.Errorf("error setting Qos, error: %v", err)
		return err
	}

	if err = r.channel.ExchangeDeclare(
		r.exchangeName,
		r.rabbitmqConsumerConfig.ExchangeOptions.Type.ToString(),
		r.rabbitmqConsumerConfig.ExchangeOptions.Durable,
		r.rabbitmqConsumerConfig.ExchangeOptions.AutoDelete,
		false,
		r.rabbitmqConsumerConfig.NoWait,
		r.rabbitmqConsumerConfig.ExchangeOptions.Args); err != nil {
		r.logger.Errorf("error declaring exchange, error: %v", err)
		return err
	}

	if _, err = r.channel.QueueDeclare(
		r.queueName,
		r.rabbitmqConsumerConfig.QueueOptions.Durable,
		r.rabbitmqConsumerConfig.QueueOptions.AutoDelete,
		r.rabbitmqConsumerConfig.QueueOptions.Exclusive,
		r.rabbitmqConsumerConfig.NoWait,
		r.rabbitmqConsumerConfig.QueueOptions.Args); err != nil {
		r.logger.Errorf("error declaring queue, error: %v", err)
		return err
	}

	if err = r.channel.QueueBind(
		r.queueName,
		r.routingKey,
		r.exchangeName,
		r.rabbitmqConsumerConfig.NoWait,
		r.rabbitmqConsumerConfig.BindingOptions.Args); err != nil {
		r.logger.Errorf("error binding queue, error: %v", err)
		return err
	}

	delivering, err := r.channel.Consume(
		r.queueName,
		r.rabbitmqConsumerConfig.ConsumerID,
		r.rabbitmqConsumerConfig.AutoAck, // When autoAck (also known as noAck) is true, the server will acknowledge deliveries to this consumer prior to writing the delivery to the network. When autoAck is true, the consumer should not call Delivery.Ack.
		r.rabbitmqConsumerConfig.QueueOptions.Exclusive,
		r.rabbitmqConsumerConfig.NoLocal,
		r.rabbitmqConsumerConfig.NoWait,
		nil,
	)
	if err != nil {
		r.logger.Errorf("error consuming messages, error: %v", err)
		return err
	}

	r.chClosedCh = r.channel.NotifyClose(make(chan *amqp.Error, 1))
	go r.handleErrorEvent()
	go r.handleReconnectionEvent(ctx)
	go r.handleChannelClosed(ctx)
	for i := range r.rabbitmqConsumerConfig.ConcurrencyLimit {
		r.logger.Infof("Processing messages on thread %d", i)
		go func() {
			for {
				delivery, ok := <-delivering
				if !ok {
					r.logger.Info("consumer connection dropped")
					return
				}
				r.handleReceived(ctx, delivery)
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

func (r *rabbitMQConsumer) handleChannelClosed(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			r.logger.Info("shutting down consumer")
			if r.channel != nil {
				if err := r.channel.Close(); err != nil {
					r.logger.Errorf("error closing channel, error: %v", err)
				}
			}
			return
		case amqErr := <-r.chClosedCh:
			if amqErr == nil {
				continue
			}
			r.logger.Errorf("AMQP Channel closed due to: %s", amqErr)
			if err := r.Start(ctx); err != nil {
				r.logger.Errorf("error restarting consumer, error: %v", err)
			}
		}
	}
}

func (r *rabbitMQConsumer) handleErrorEvent() {
	// close delivery
	for errEvent := range r.connection.ErrorEvent() {
		r.logger.Errorf("rabbitmq connection error: %v", errEvent)
		if r.connection == nil {
			continue
		}
		if r.connection.IsClosed() {
			continue
		}
		if r.channel == nil {
			continue
		}
		if err := r.channel.Cancel(r.rabbitmqConsumerConfig.ConsumerID, false); err != nil {
			r.logger.Errorf("error closing channel, error: %v", err)
		}
	}
}

func (r *rabbitMQConsumer) handleReconnectionEvent(ctx context.Context) {
	for range r.connection.ReconnectedEvent() {
		// close the channel
		if r.channel != nil {
			if err := r.channel.Close(); err != nil {
				r.logger.Errorf("error closing channel, error: %v", err)
			}
		}
		r.logger.Info("restarting consumer")
		if err := r.Start(ctx); err != nil {
			r.logger.Error("restarting consumer failed with error: %v", err)
		}
		r.logger.Info("restarting consumer finished successfully")
	}
}

func (r *rabbitMQConsumer) handleReceived(ctx context.Context, delivery amqp.Delivery) {
	// for ensuring our handlers execute completely after shutdown
	r.deliveryRoutines <- struct{}{}
	defer func() { <-r.deliveryRoutines }()

	consumeContext, err := r.createMessageConsumeContext(delivery)
	if err != nil {
		r.logger.Errorf("error creating message consume context, error: %v", err)
		return
	}

	if r.rabbitmqConsumerConfig.AutoAck {
		r.handle(ctx, nil, nil, consumeContext)
		return
	}

	customizeAck := func() {
		if err := delivery.Ack(false); err != nil {
			r.logger.Errorf(
				"error sending ACK to RabbitMQ consumer, error: %v, correlation id: %s, message id: %s, message type: %s",
				err,
				consumeContext.CorrelationID(),
				consumeContext.MessageID(),
				consumeContext.Type(),
			)
			return
		}
		for _, cousumedFunc := range r.consumedFuncs {
			if cousumedFunc == nil {
				continue
			}
			cousumedFunc(consumeContext.Message())
		}
	}
	// requeue with retry count
	customizeNack := func(err error) {
		retryCount, _ := delivery.Headers[enums.DeliveryHeaderRetryCount.ToString()].(int32)
		if retryCount >= int32(r.rabbitmqConfig.RetryAttempts) {
			// publish message to DLQ
			r.publichToDLQ(ctx, &delivery, err)
		} else {
			// re queue with retry count
			delivery.Headers[enums.DeliveryHeaderRetryCount.ToString()] = retryCount + 1
			r.requeueMessage(ctx, &delivery)
		}
	}
	r.handle(ctx, &customizeAck, &customizeNack, consumeContext)
}

func (r *rabbitMQConsumer) publichToDLQ(ctx context.Context, delivery *amqp.Delivery, err error) {
	msg := rabbitmq.ConvertDeliveryToPublishing(delivery, true)
	msg = *rabbitmq.NewPublishingHeaderSetter(&msg).
		SetError(err).
		SetRetryCount(int(delivery.Headers[enums.DeliveryHeaderRetryCount.ToString()].(int32))).
		SetQueue(r.queueName).
		SetExchange(r.exchangeName).
		SetRoutingKey(r.routingKey).
		Build()

	r.logger.Infof(
		"Max retries reached. Moving to DLQ: id: %s, message id: %s, message type: %s",
		delivery.CorrelationId,
		delivery.MessageId,
		delivery.Type,
	)

	if err := r.channel.PublishWithContext(
		ctx,
		r.rabbitmqConfig.DeadLetterExchange,
		r.rabbitmqConfig.DeadLetterRoutingKey,
		false,
		false,
		msg,
	); err != nil {
		r.logger.Errorf("error publishing message to DLQ, error: %v", err)
	}
	if err := delivery.Ack(false); err != nil {
		r.logger.Errorf(
			"error sending dlq ACK to RabbitMQ consumer, error: %v, correlation id: %s, message id: %s, message type: %s",
			err,
			delivery.CorrelationId,
			delivery.MessageId,
			delivery.Type,
		)
	}
}

func (r *rabbitMQConsumer) requeueMessage(ctx context.Context, delivery *amqp.Delivery) {
	msg := rabbitmq.ConvertDeliveryToPublishing(delivery, true)
	if err := r.channel.PublishWithContext(
		ctx,
		r.exchangeName,
		r.routingKey,
		false,
		false,
		msg,
	); err != nil {
		r.logger.Errorf("error requeueing message to RabbitMQ consumer, error: %v", err)
	}
	if err := delivery.Ack(false); err != nil {
		r.logger.Errorf(
			"error sending requeue ACK to RabbitMQ consumer, error: %v, correlation id: %s, message id: %s, message type: %s",
			err,
			delivery.CorrelationId,
			delivery.MessageId,
			delivery.Type,
		)
	}
}

func (r *rabbitMQConsumer) handle(
	ctx context.Context,
	customizeAck *func(),
	customizeNack *func(err error),
	messageConsumeContext types.MessageConsumeContext,
) {
	for _, handler := range r.rabbitmqConsumerConfig.Handlers {
		err := r.handleWithRetry(ctx, handler, messageConsumeContext)
		if err != nil {
			r.logger.Errorf("error handling consume message, prepare for nacking message, error: %v", err)
			if customizeNack != nil {
				(*customizeNack)(err)
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
	r.logger.Infof("handling message, correlation id: %s, message id: %s, message type: %s",
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
