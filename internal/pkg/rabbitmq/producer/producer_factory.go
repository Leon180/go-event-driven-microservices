package producer

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/connect"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
)

type ProducerFactory interface {
	CreateProducer(
		rabbitmqProducersConfigs []*RabbitMQProducerConfig,
		producedFuncs []producer.ProducedFunc,
	) producer.Producer
}

func NewProducerFactory(
	rabbitmqConfig *rabbitmq.RabbitMQConfig,
	connection connect.AMQPConnection,
	eventSerializer serializers.MessageSerializer,
	logger loggers.Logger,
	uuidGenerator uuid.UUIDGenerator,
) ProducerFactory {
	return &producerFactory{
		uuidGenerator:   uuidGenerator,
		rabbitmqConfig:  rabbitmqConfig,
		logger:          logger,
		connection:      connection,
		eventSerializer: eventSerializer,
	}
}

type producerFactory struct {
	connection      connect.AMQPConnection
	logger          loggers.Logger
	eventSerializer serializers.MessageSerializer
	rabbitmqConfig  *rabbitmq.RabbitMQConfig
	uuidGenerator   uuid.UUIDGenerator
}

func (p *producerFactory) CreateProducer(
	rabbitmqProducersConfigs []*RabbitMQProducerConfig,
	producedFuncs []producer.ProducedFunc,
) producer.Producer {
	return NewRabbitMQProducer(
		p.rabbitmqConfig,
		p.connection,
		rabbitmqProducersConfigs,
		p.logger,
		p.uuidGenerator,
		p.eventSerializer,
		producedFuncs,
	)
}
