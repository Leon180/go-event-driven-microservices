package operators

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqproducer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/producer"

	"github.com/samber/lo"
)

type RabbitMQOperatorsConfigBuilderFunc func(builder RabbitMQOperatorsConfigBuilder)

type RabbitMQOperatorsConfigBuilder interface {
	AddProducer(
		producerMessageType types.Message,
		producerBuilderFunc rabbitmqproducer.RabbitMQProducerConfigBuilderFunc,
	) RabbitMQOperatorsConfigBuilder
	AddConsumer(
		consumerMessageType types.Message,
		consumerBuilderFunc rabbitmqconsumer.RabbitMQConsumerConfigBuilderFunc,
	) RabbitMQOperatorsConfigBuilder
	Build() *RabbitMQOperatorsConfig
}

func NewRabbitMQOperatorsConfigBuilder() RabbitMQOperatorsConfigBuilder {
	return &rabbitMQOperatorsConfigBuilder{
		rabbitMQOperatorsConfig: &RabbitMQOperatorsConfig{},
	}
}

type rabbitMQOperatorsConfigBuilder struct {
	rabbitMQOperatorsConfig *RabbitMQOperatorsConfig
	consumerBuilders        []rabbitmqconsumer.RabbitMQConsumerConfigBuilder
	producerBuilders        []rabbitmqproducer.RabbitMQProducerConfigBuilder
}

func (r *rabbitMQOperatorsConfigBuilder) AddProducer(
	producerMessageType types.Message,
	producerBuilderFunc rabbitmqproducer.RabbitMQProducerConfigBuilderFunc,
) RabbitMQOperatorsConfigBuilder {
	builder := rabbitmqproducer.NewRabbitMQProducerConfigBuilder(producerMessageType)
	if producerBuilderFunc != nil {
		producerBuilderFunc(builder)
	}
	r.producerBuilders = append(r.producerBuilders, builder)
	return r
}

func (r *rabbitMQOperatorsConfigBuilder) AddConsumer(
	consumerMessageType types.Message,
	consumerBuilderFunc rabbitmqconsumer.RabbitMQConsumerConfigBuilderFunc,
) RabbitMQOperatorsConfigBuilder {
	builder := rabbitmqconsumer.NewRabbitMQConsumerConfigBuilder(consumerMessageType)
	if consumerBuilderFunc != nil {
		consumerBuilderFunc(builder)
	}
	r.consumerBuilders = append(r.consumerBuilders, builder)
	return r
}

func (r *rabbitMQOperatorsConfigBuilder) Build() *RabbitMQOperatorsConfig {
	consumersConfigs := lo.Map(
		r.consumerBuilders,
		func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder, index int) *rabbitmqconsumer.RabbitMQConsumerConfig {
			return builder.Build()
		},
	)

	producersConfigs := lo.Map(
		r.producerBuilders,
		func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder, index int) *rabbitmqproducer.RabbitMQProducerConfig {
			return builder.Build()
		},
	)

	r.rabbitMQOperatorsConfig.ConsumersConfigs = consumersConfigs
	r.rabbitMQOperatorsConfig.ProducersConfigs = producersConfigs
	return r.rabbitMQOperatorsConfig
}
