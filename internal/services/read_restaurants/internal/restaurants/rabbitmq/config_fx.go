package restaurantsrabbitmq

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqoperators "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/operators"
	createbookcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/commands"
	createbookevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/events"
	createrestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/commands"
	createrestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/events"
	deletebookcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/commands"
	deletebookevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/events"
	deleterestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/commands"
	deleterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/events"
	restorerestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/commands"
	restorerestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/events"
	updaterestaurantcommands "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/commands"
	updaterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/events"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"restaurantsRabbitMQProvideModule",
	fx.Provide(
		NewReadRestaurantsRabbitMQOperatorsConfigBuilderFunc,
		fx.Annotate(
			NewReadRestaurantsRabbitMQOperatorsConfigBuilderFunc,
			fx.As(new(rabbitmqoperators.RabbitMQOperatorsConfigBuilderFunc)),
		),
	),
)

func NewReadRestaurantsRabbitMQOperatorsConfigBuilderFunc(
	logger loggers.Logger,
	createBookCommand createbookcommands.CreateBookHandler,
	createRestaurantCommand createrestaurantcommands.CreateRestaurantHandler,
	deleteBookCommand deletebookcommands.DeleteBookHandler,
	deleteRestaurantCommand deleterestaurantcommands.DeleteRestaurantHandler,
	restoreRestaurantCommand restorerestaurantcommands.RestoreRestaurantHandler,
	updateRestaurantCommand updaterestaurantcommands.UpdateRestaurantHandler,
) rabbitmqoperators.RabbitMQOperatorsConfigBuilderFunc {
	return func(builder rabbitmqoperators.RabbitMQOperatorsConfigBuilder) {
		builder.
			AddConsumer(
				createbookevents.CreateBook{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						createbookevents.NewCreateBookHandler(logger, createBookCommand),
					)
				}).
			AddConsumer(
				createrestaurantevents.CreateRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						createrestaurantevents.NewCreateRestaurantHandler(logger, createRestaurantCommand),
					)
				}).
			AddConsumer(
				deletebookevents.DeleteBook{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						deletebookevents.NewDeleteBookHandler(logger, deleteBookCommand),
					)
				}).
			AddConsumer(
				deleterestaurantevents.DeleteRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						deleterestaurantevents.NewDeleteRestaurantHandler(logger, deleteRestaurantCommand),
					)
				}).
			AddConsumer(
				restorerestaurantevents.RestoreRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						restorerestaurantevents.NewRestoreRestaurantHandler(logger, restoreRestaurantCommand),
					)
				}).
			AddConsumer(
				updaterestaurantevents.UpdateRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						updaterestaurantevents.NewUpdateRestaurantHandler(logger, updateRestaurantCommand),
					)
				})
	}
}
