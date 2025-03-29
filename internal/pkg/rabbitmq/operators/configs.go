package operators

import (
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqproducer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/producer"
)

type RabbitMQOperatorsConfig struct {
	ProducersConfigs []*rabbitmqproducer.RabbitMQProducerConfig
	ConsumersConfigs []*rabbitmqconsumer.RabbitMQConsumerConfig
}
