package producer

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	messageheader "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/message_header"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/connect"
	customizereflect "github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
)

func NewRabbitMQProducer(
	cfg *rabbitmq.RabbitMQConfig,
	connection connect.AMQPConnection,
	rabbitmqProducersConfigs []*RabbitMQProducerConfig,
	logger loggers.Logger,
	uuidGenerator uuid.UUIDGenerator,
	eventSerializer serializers.MessageSerializer,
	producedFuncs []producer.ProducedFunc,
) producer.Producer {
	return &rabbitMQProducer{
		logger:            logger,
		uuidGenerator:     uuidGenerator,
		rabbitmqOptions:   cfg,
		connection:        connection,
		messageSerializer: eventSerializer,
		producedFuncs:     producedFuncs,
		producersConfigurations: lo.SliceToMap(
			rabbitmqProducersConfigs,
			func(item *RabbitMQProducerConfig) (reflect.Type, *RabbitMQProducerConfig) {
				return item.ProducerMessageType, item
			},
		),
	}
}

type rabbitMQProducer struct {
	logger                  loggers.Logger
	uuidGenerator           uuid.UUIDGenerator
	rabbitmqOptions         *rabbitmq.RabbitMQConfig
	connection              connect.AMQPConnection
	messageSerializer       serializers.MessageSerializer
	producersConfigurations map[reflect.Type]*RabbitMQProducerConfig
	producedFuncs           []producer.ProducedFunc
}

func (r *rabbitMQProducer) PublishMessage(
	ctx context.Context,
	message types.Message,
	meta metadatas.Metadata,
	topicOrExchangeName *string,
) error {
	producerConfiguration := r.getProducerConfigurationByMessage(message)

	if producerConfiguration == nil {
		producerConfiguration = NewDefaultRabbitMQProducerConfig(message)
	}

	exchangeName := ""
	if topicOrExchangeName != nil {
		exchangeName = *topicOrExchangeName
	}
	if exchangeName == "" {
		exchangeName = producerConfiguration.ExchangeOptions.Name
	}
	if exchangeName == "" {
		exchangeName = customizereflect.GetAnysTypeName(message)
	}

	routingKey := producerConfiguration.BindingOptions.RoutingKey
	if routingKey == "" {
		routingKey = customizereflect.GetAnysTypeName(message)
	}

	if meta == nil {
		meta = metadatas.NewMetadata()
		meta.Set(enums.DeliveryHeaderRetryCount.ToString(), 0)
	}

	meta = r.getMetadata(message, meta)

	serialization, err := r.messageSerializer.Serialize(message)
	if err != nil {
		return err
	}

	channel, err := r.connection.NewChannel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = r.ensureExchange(producerConfiguration, channel, exchangeName)
	if err != nil {
		return err
	}

	if err := channel.Confirm(false); err != nil {
		return err
	}

	confirms := make(chan amqp.Confirmation)
	channel.NotifyPublish(confirms)

	msg := amqp.Publishing{
		CorrelationId:   messageheader.GetCorrelationID(meta),
		MessageId:       message.ID(),
		Timestamp:       time.Now(),
		Headers:         metadatas.MetadataToMap(meta),
		Type:            message.Type(),
		ContentType:     serialization.ContentType.ToString(),
		Body:            serialization.Data,
		DeliveryMode:    producerConfiguration.DeliveryMode,
		Expiration:      producerConfiguration.Expiration,
		AppId:           producerConfiguration.AppId,
		Priority:        producerConfiguration.Priority,
		ReplyTo:         producerConfiguration.ReplyTo,
		ContentEncoding: producerConfiguration.ContentEncoding,
	}

	err = channel.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		true,
		false,
		msg,
	)
	if err != nil {
		return err
	}

	r.logger.Infof(
		"message published, correlation id: %s, message id: %s, message type: %s, message name: %s, message content type: %s, message headers: %v, message body: %v",
		messageheader.GetCorrelationID(meta),
		message.ID(),
		message.Type(),
		messageheader.GetMessageName(meta),
		messageheader.GetContentType(meta),
	)

	if confirmed := <-confirms; !confirmed.Ack {
		return errors.New("ack not confirmed")
	}

	r.logger.Infof(
		"message ack confirmed, correlation id: %s, message id: %s, message type: %s, message name: %s, message content type: %s, message headers: %v, message body: %v",
		messageheader.GetCorrelationID(meta),
		message.ID(),
		message.Type(),
		messageheader.GetMessageName(meta),
		messageheader.GetContentType(meta),
	)

	for _, producedFunc := range r.producedFuncs {
		if producedFunc != nil {
			producedFunc(message)
		}
	}

	return nil
}

func (r *rabbitMQProducer) getProducerConfigurationByMessage(
	message types.Message,
) *RabbitMQProducerConfig {
	messageType := customizereflect.GetAnysType(message)
	return r.producersConfigurations[messageType]
}

func (r *rabbitMQProducer) getMetadata(
	message types.Message,
	meta metadatas.Metadata,
) metadatas.Metadata {
	// just message type name not full type name because in other side package name for type could be different
	messageheader.SetMessageType(meta, message.Type())
	messageheader.SetContentType(meta, r.messageSerializer.ContentType().ToString())

	if messageheader.GetMessageID(meta) == "" {
		messageheader.SetMessageID(meta, message.ID())
	}

	if messageheader.GetTimeStamp(meta) == *new(time.Time) {
		messageheader.SetTimeStamp(meta, message.TimeStamp())
	}

	if messageheader.GetCorrelationID(meta) == "" {
		cid := r.uuidGenerator.GenerateUUID()
		messageheader.SetCorrelationID(meta, cid)
	}
	messageheader.SetMessageName(meta, customizereflect.GetAnysTypeName(message))

	return meta
}

func (r *rabbitMQProducer) ensureExchange(
	producersConfigurations *RabbitMQProducerConfig,
	channel *amqp091.Channel,
	exchangeName string,
) error {
	return channel.ExchangeDeclare(
		exchangeName,
		string(producersConfigurations.ExchangeOptions.Type),
		producersConfigurations.ExchangeOptions.Durable,
		producersConfigurations.ExchangeOptions.AutoDelete,
		false,
		false,
		producersConfigurations.ExchangeOptions.Args,
	)
}

func (r *rabbitMQProducer) Produced(producedFuncs ...producer.ProducedFunc) {
	r.producedFuncs = append(r.producedFuncs, producedFuncs...)
}
