package restaurantsrabbitmq

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqoperators "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/operators"
	createbookevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/events"
	createbookservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/services"
	createrestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/events"
	createrestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/services"
	deletebookevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/events"
	deletebookservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/services"
	deleterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/events"
	deleterestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
	restorerestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/events"
	restorerestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/services"
	updaterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/events"
	updaterestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/services"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"restaurantsRabbitMQProvideModule",
	fx.Provide(
		NewReadRestaurantsRabbitMQOperatorsConfigBuilderFunc,
	),
)

func NewReadRestaurantsRabbitMQOperatorsConfigBuilderFunc(
	logger loggers.Logger,
	createBookService createbookservices.CreateBookHandler,
	createRestaurantService createrestaurantservices.CreateRestaurantHandler,
	deleteBookService deletebookservices.DeleteBookHandler,
	deleteRestaurantService deleterestaurantservices.DeleteRestaurantHandler,
	restoreRestaurantService restorerestaurantservices.RestoreRestaurantHandler,
	updateRestaurantService updaterestaurantservices.UpdateRestaurantHandler,
) rabbitmqoperators.RabbitMQOperatorsConfigBuilderFunc {
	return func(builder rabbitmqoperators.RabbitMQOperatorsConfigBuilder) {
		builder.
			AddConsumer(
				createbookevents.CreateBook{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						createbookevents.NewCreateBookHandler(logger, createBookService),
					)
				}).
			AddConsumer(
				createrestaurantevents.CreateRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						createrestaurantevents.NewCreateRestaurantHandler(logger, createRestaurantService),
					)
				}).
			AddConsumer(
				deletebookevents.DeleteBook{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						deletebookevents.NewDeleteBookHandler(logger, deleteBookService),
					)
				}).
			AddConsumer(
				deleterestaurantevents.DeleteRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						deleterestaurantevents.NewDeleteRestaurantHandler(logger, deleteRestaurantService),
					)
				}).
			AddConsumer(
				restorerestaurantevents.RestoreRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						restorerestaurantevents.NewRestoreRestaurantHandler(logger, restoreRestaurantService),
					)
				}).
			AddConsumer(
				updaterestaurantevents.UpdateRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						updaterestaurantevents.NewUpdateRestaurantHandler(logger, updateRestaurantService),
					)
				})
	}
}
