package restaurantsrabbitmq

import (
	rabbitmqoperators "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/operators"
	rabbitmqproducer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/producer"
	createrestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_restaurant/events"
	deleterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/delete_restaurant/events"
	restorerestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/restore_restaurant/events"
	synccategoriesevents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/sync_categories/events"
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
				createrestaurantevents.CreateRestaurant{},
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
			).
			AddProducer(
				synccategoriesevents.SyncCategories{},
				func(builder rabbitmqproducer.RabbitMQProducerConfigBuilder) {
				},
			)
	}
}
