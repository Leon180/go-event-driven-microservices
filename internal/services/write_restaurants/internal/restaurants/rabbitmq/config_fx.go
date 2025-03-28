package restaurantsrabbitmq

import (
	rabbitmqoperators "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/operators"
	rabbitmqproducer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/producer"
	createbookevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_book/events"
	createrestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_restaurant/events"
	deletebookevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_book/events"
	deleterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/events"
	restorerestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/events"
	updaterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/update_restaurant/events"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"restaurantsRabbitMQProvideModule",
	fx.Provide(
		NewWriteRestaurantsRabbitMQOperatorsConfigBuilderFunc,
	),
)

func NewWriteRestaurantsRabbitMQOperatorsConfigBuilderFunc() rabbitmqoperators.RabbitMQOperatorsConfigBuilderFunc {
	return func(builder rabbitmqoperators.RabbitMQOperatorsConfigBuilder) {
		builder.
			AddProducer(
				createbookevents.CreateBook{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			).
			AddProducer(
				createrestaurantevents.CreateRestaurant{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			).
			AddProducer(
				deletebookevents.DeleteBook{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			).
			AddProducer(
				deleterestaurantevents.DeleteRestaurant{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			).
			AddProducer(
				restorerestaurantevents.RestoreRestaurant{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			).
			AddProducer(
				updaterestaurantevents.UpdateRestaurant{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			)
	}
}
