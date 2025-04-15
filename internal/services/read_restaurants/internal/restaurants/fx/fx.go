package accountsfx

import (
	customizeginfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/fx"
	customizegrpcfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc/fx"
	environmentsfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/environments/fx"
	loggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/fx"
	messagingfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/fx"
	rabbitmqfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq/fx"
	redisdbfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/redisdb/fx"
	contextloggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers/fx"
	uuidfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/fx"
	appconfigsfx "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/configs/fx"
	grpcfx "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/fx"
	featuresfx "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/fx"
	mongodbfx "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/mongodb/fx"
	restaurantsrabbitmq "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/rabbitmq"
	repositoriesfx "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories/fx"
	"go.uber.org/fx"
)

var RestaurantsConfiguratorModule = fx.Module(
	"restaurantsConfiguratorFx",

	ProvideModule,
	InvokeModule,
)

var ProvideModule = fx.Module(
	"restaurantsProvideFx",

	// environments
	environmentsfx.ProvideModule,

	// uuid generators
	uuidfx.ProvideModule,

	// loggers
	loggersfx.ProvideModule,
	contextloggersfx.ProvideModule,

	// db
	redisdbfx.ProvideModule,
	mongodbfx.ProvideModule,

	// messaging serializers
	messagingfx.ProvideModule,

	// rabbitmq
	rabbitmqfx.ProvideModule,
	restaurantsrabbitmq.ProvideModule,

	// app configs
	appconfigsfx.ProvideModule,

	// repositories
	repositoriesfx.ProvideModule,

	// features
	featuresfx.ProvideModule,

	// grpc server
	customizegrpcfx.ProvideModule,
	grpcfx.ProvideModule,

	// gin server
	customizeginfx.ProvideModule,
)

var InvokeModule = fx.Module(
	"restaurantsInvokeFx",

	// messaging
	messagingfx.InvokeModule,

	// rabbitmq
	rabbitmqfx.InvokeModule,

	// grpc server
	customizegrpcfx.InvokeModule,

	// start server
	customizeginfx.InvokeModule,
)
