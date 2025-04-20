package restaurantsrabbitmq

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	serializers "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	types "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/types"
	rabbitmq "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	rabbitmqconsumer "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/consumer"
	rabbitmqoperators "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/operators"
	createfailedmessageevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/events"
	createfailedmessageservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/services"
	createrestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/events"
	createrestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/services"
	deleterestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/events"
	deleterestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
	restorerestaurantevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/events"
	restorerestaurantservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/services"
	synccategoriesevents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/sync_categories/events"
	synccategorieservices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/sync_categories/services"
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
	createRestaurantService createrestaurantservices.CreateRestaurantHandler,
	deleteRestaurantService deleterestaurantservices.DeleteRestaurantHandler,
	restoreRestaurantService restorerestaurantservices.RestoreRestaurantHandler,
	updateRestaurantService updaterestaurantservices.UpdateRestaurantHandler,
	syncCategoriesService synccategorieservices.SyncCategoriesHandler,
	createFailedMessageService createfailedmessageservices.CreateFailedMessageHandler,
	messageSerializer serializers.MessageSerializer,
	rabbitMQConfig *rabbitmq.RabbitMQConfig,
) rabbitmqoperators.RabbitMQOperatorsConfigBuilderFunc {
	return func(builder rabbitmqoperators.RabbitMQOperatorsConfigBuilder) {
		builder.
			AddConsumer(
				createrestaurantevents.CreateRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						createrestaurantevents.NewCreateRestaurantHandler(logger, createRestaurantService),
					)
				},
			).
			AddConsumer(
				deleterestaurantevents.DeleteRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						deleterestaurantevents.NewDeleteRestaurantHandler(logger, deleteRestaurantService),
					)
				},
			).
			AddConsumer(
				restorerestaurantevents.RestoreRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						restorerestaurantevents.NewRestoreRestaurantHandler(logger, restoreRestaurantService),
					)
				},
			).
			AddConsumer(
				updaterestaurantevents.UpdateRestaurant{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						updaterestaurantevents.NewUpdateRestaurantHandler(logger, updateRestaurantService),
					)
				},
			).
			AddConsumer(
				synccategoriesevents.SyncCategories{},
				func(builder rabbitmqconsumer.RabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						synccategoriesevents.NewSyncCategoriesHandler(logger, syncCategoriesService),
					)
				},
			).
			SetDeadLetterConsumer(
				[]types.Message{
					createrestaurantevents.CreateRestaurant{},
					deleterestaurantevents.DeleteRestaurant{},
					restorerestaurantevents.RestoreRestaurant{},
					updaterestaurantevents.UpdateRestaurant{},
					synccategoriesevents.SyncCategories{},
				},
				func(builder rabbitmqconsumer.DeadLetterRabbitMQConsumerConfigBuilder) {
					builder.SetHandlers(
						createfailedmessageevents.NewCreateFailedMessageHandler(
							logger,
							createFailedMessageService,
							messageSerializer,
						),
					)
				},
				rabbitMQConfig,
			)
	}
}
