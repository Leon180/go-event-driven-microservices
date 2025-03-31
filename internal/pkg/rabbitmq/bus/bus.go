package bus

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/bus"
	consumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/consumer"
	producer "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/producer"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/metadatas"
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqoperators "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/operators"
	rabbitmqproducer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/producer"
	customizereflect "github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type RabbitMQBus interface {
	bus.Bus
	rabbitmqconsumer.RabbitMQConsumerConnector
}

func NewRabbitMQBus(
	logger loggers.Logger,
	consumerFactory rabbitmqconsumer.ConsumerFactory,
	producerFactory rabbitmqproducer.ProducerFactory,
	rabbitmqBuilderFunc rabbitmqoperators.RabbitMQOperatorsConfigBuilderFunc,
) (RabbitMQBus, error) {
	builder := rabbitmqoperators.NewRabbitMQOperatorsConfigBuilder()

	if rabbitmqBuilderFunc != nil {
		rabbitmqBuilderFunc(builder)
	}

	rabbitmqOperatorsConfig := builder.Build()
	rabbitBus := &rabbitmqBus{
		logger:                         logger,
		rabbitmqOperatorsConfig:        rabbitmqOperatorsConfig,
		consumerFactory:                consumerFactory,
		producerFactory:                producerFactory,
		rabbitmqOperatorsConfigBuilder: builder,
		messageTypeConsumers:           map[reflect.Type][]consumer.Consumer{},
	}

	if rabbitBus.rabbitmqOperatorsConfig.DeadLetterConsumersConfigs != nil {
		rabbitBus.deadLetterConsumer = consumerFactory.CreateDeadLetterConsumer(
			&rabbitBus.rabbitmqOperatorsConfig.DeadLetterConsumersConfigs.RabbitMQConsumerConfig,
			[]consumer.ConsumedFunc{func(message types.Message) {
				for _, consumedFunc := range rabbitBus.deadLetterConsumedFuncs {
					if consumedFunc != nil {
						consumedFunc(message)
					}
				}
			}},
		)
	}

	for _, consumerConfig := range rabbitBus.rabbitmqOperatorsConfig.ConsumersConfigs {
		mqConsumer := consumerFactory.CreateConsumer(
			consumerConfig,
			[]consumer.ConsumedFunc{func(message types.Message) {
				for _, consumedFunc := range rabbitBus.consumedFuncs {
					if consumedFunc != nil {
						consumedFunc(message)
					}
				}
			}},
		)
		rabbitBus.messageTypeConsumers[consumerConfig.ConsumerMessageType] = append(
			rabbitBus.messageTypeConsumers[consumerConfig.ConsumerMessageType],
			mqConsumer,
		)
	}

	rabbitBus.producer = producerFactory.CreateProducer(
		rabbitBus.rabbitmqOperatorsConfig.ProducersConfigs,
		[]producer.ProducedFunc{func(message types.Message) {
			for _, producedFunc := range rabbitBus.producedFuncs {
				if producedFunc != nil {
					producedFunc(message)
				}
			}
		}},
	)

	return rabbitBus, nil
}

type rabbitmqBus struct {
	messageTypeConsumers           map[reflect.Type][]consumer.Consumer
	deadLetterConsumer             consumer.Consumer
	producer                       producer.Producer
	rabbitmqOperatorsConfig        *rabbitmqoperators.RabbitMQOperatorsConfig
	rabbitmqOperatorsConfigBuilder rabbitmqoperators.RabbitMQOperatorsConfigBuilder
	logger                         loggers.Logger
	consumerFactory                rabbitmqconsumer.ConsumerFactory
	producerFactory                rabbitmqproducer.ProducerFactory
	consumedFuncs                  []consumer.ConsumedFunc
	producedFuncs                  []producer.ProducedFunc
	deadLetterConsumedFuncs        []consumer.ConsumedFunc
}

func (r *rabbitmqBus) Consumed(consumedFuncs ...consumer.ConsumedFunc) {
	r.consumedFuncs = append(r.consumedFuncs, consumedFuncs...)
}

func (r *rabbitmqBus) Produced(producedFuncs ...producer.ProducedFunc) {
	r.producedFuncs = append(r.producedFuncs, producedFuncs...)
}

func (r *rabbitmqBus) ConnectConsumer(message types.Message, consumer consumer.Consumer) {
	typ := customizereflect.GetAnysType(message)
	r.messageTypeConsumers[typ] = append(r.messageTypeConsumers[typ], consumer)
}

func (r *rabbitmqBus) ConnectRabbitMQConsumer(
	message types.Message,
	consumerBuilderFunc rabbitmqconsumer.RabbitMQConsumerConfigBuilderFunc,
) {
	builder := rabbitmqconsumer.NewRabbitMQConsumerConfigBuilder(message)
	if consumerBuilderFunc != nil {
		consumerBuilderFunc(builder)
	}
	consumerConfig := builder.Build()
	mqConsumer := r.consumerFactory.CreateConsumer(
		consumerConfig,
		[]consumer.ConsumedFunc{func(message types.Message) {
			for _, consumedFunc := range r.consumedFuncs {
				if consumedFunc != nil {
					consumedFunc(message)
				}
			}
		}},
	)

	typ := customizereflect.GetAnysType(message)
	r.messageTypeConsumers[typ] = append(r.messageTypeConsumers[typ], mqConsumer)
}

func (r *rabbitmqBus) ConnectConsumerHandler(message types.Message, consumerHandler consumer.ConsumerHandler) {
	typ := customizereflect.GetAnysType(message)
	consumers := r.messageTypeConsumers[typ]
	for _, c := range consumers {
		c.BindHandler(consumerHandler)
	}
	consumerBuilder := rabbitmqconsumer.NewRabbitMQConsumerConfigBuilder(message)
	consumerBuilder.SetHandlers(consumerHandler)
	consumerConfig := consumerBuilder.Build()
	consumer := r.consumerFactory.CreateConsumer(
		consumerConfig,
		[]consumer.ConsumedFunc{func(message types.Message) {
			for _, consumedFunc := range r.consumedFuncs {
				if consumedFunc != nil {
					consumedFunc(message)
				}
			}
		}},
	)
	r.messageTypeConsumers[typ] = append(r.messageTypeConsumers[typ], consumer)
}

func (r *rabbitmqBus) Start(ctx context.Context) error {
	r.logger.Infof("rabbitmq is running on host: %s", r.consumerFactory.Connection().Raw().LocalAddr().String())

	for messageType, consumers := range r.messageTypeConsumers {
		name := customizereflect.GetTypeNameFromType(messageType)
		r.logger.Info(fmt.Sprintf("consuming message type %s", name))
		for _, consumer := range consumers {
			err := consumer.Start(ctx)
			r.logger.Info("start consumer: ", consumer.Name())
			if err != nil {
				r.logger.Error("error in consumer: ", consumer.Name(), err)
			}
		}
	}

	if r.deadLetterConsumer != nil {
		err := r.deadLetterConsumer.Start(ctx)
		r.logger.Info("start dead letter consumer: ", r.deadLetterConsumer.Name())
		if err != nil {
			r.logger.Error("error in dead letter consumer: ", r.deadLetterConsumer.Name(), err)
		}
	}

	return nil
}

func (r *rabbitmqBus) Stop() error {
	waitGroup := sync.WaitGroup{}
	for _, consumers := range r.messageTypeConsumers {
		for _, c := range consumers {
			waitGroup.Add(1)
			go func(c consumer.Consumer) {
				defer waitGroup.Done()
				err := c.Stop()
				if err != nil {
					r.logger.Error("error in the unconsuming")
				}
			}(c)
		}
	}
	waitGroup.Wait()

	return nil
}

func (r *rabbitmqBus) PublishMessage(
	ctx context.Context,
	message types.Message,
	meta metadatas.Metadata,
	topicOrExchangeName *string,
) error {
	return r.producer.PublishMessage(ctx, message, meta, topicOrExchangeName)
}
