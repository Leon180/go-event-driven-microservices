package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createFailedMessageEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/events"
	createFailedMessageServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/services"
	createRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/events"
	createRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/services"
	deleteRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/events"
	deleteRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
	getRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/gin_endpoints"
	getRestaurantGRPCService "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/grpc"
	getRestaurantQueries "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/queries"
	listCategoriesGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/gin_endpoints"
	listCategoriesQueries "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/queries"
	restoreRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/events"
	restoreRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/services"
	searchRestaurantsGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/gin_endpoints"
	searchRestaurantsGRPCService "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/grpc"
	searchRestaurantsQueries "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/queries"
	syncCategoriesEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/sync_categories/events"
	syncCategoriesServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/sync_categories/services"
	updateRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/events"
	updateRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/update_restaurant/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the restaurants features services, commands and endpoints
var ProvideModule = fx.Module(
	"restaurantsFeaturesProvideFx",

	// services
	fx.Provide(
		createRestaurantServices.NewCreateRestaurantHandler,
		deleteRestaurantServices.NewDeleteRestaurantHandler,
		restoreRestaurantServices.NewRestoreRestaurantHandler,
		updateRestaurantServices.NewUpdateRestaurantHandler,
		syncCategoriesServices.NewSyncCategoriesHandler,
		createFailedMessageServices.NewCreateFailedMessageHandler,
	),

	// queries
	fx.Provide(
		getRestaurantQueries.NewGetRestaurantHandler,
		listCategoriesQueries.NewListCategoriesHandler,
		searchRestaurantsQueries.NewSearchRestaurantsHandler,
	),

	// events
	fx.Provide(
		createRestaurantEvents.NewCreateRestaurantHandler,
		deleteRestaurantEvents.NewDeleteRestaurantHandler,
		restoreRestaurantEvents.NewRestoreRestaurantHandler,
		updateRestaurantEvents.NewUpdateRestaurantHandler,
		syncCategoriesEvents.NewSyncCategoriesHandler,
		createFailedMessageEvents.NewCreateFailedMessageHandler,
	),

	// grpc service register
	fx.Provide(
		fxTagGRPCServiceRegisters(
			getRestaurantGRPCService.NewGRPCServiceRegister,
			searchRestaurantsGRPCService.NewGRPCServiceRegister,
		)...,
	),

	// endpoints
	fx.Provide(
		fxTagGroupEndpoints(
			getRestaurantGinEndpoints.NewGetRestaurant,
			searchRestaurantsGinEndpoints.NewSearchRestaurants,
			listCategoriesGinEndpoints.NewListCategories,
		)...,
	),
)

func fxTagGRPCServiceRegisters(handlers ...any) []any {
	return lo.Map(handlers, func(handler any, _ int) any {
		return fxTagGRPCServiceRegister(handler)
	})
}

func fxTagGRPCServiceRegister(handler any) any {
	return fx.Annotate(
		handler,
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupGRPCServiceRegister.ToString())),
	)
}

// fxTagEndpoints will tag the endpoints with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagGroupEndpoints(handlers ...any) []any {
	return lo.Map(handlers, func(handler any, _ int) any {
		return fxTagGroupEndpoint(handler)
	})
}

// fxTagEndpoint will tag the endpoint with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagGroupEndpoint(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(customizegin.Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupEndpoints.ToString())),
	)
}
