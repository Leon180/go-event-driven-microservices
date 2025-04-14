package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/events"
	createBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_book/services"
	createFailedMessageEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/events"
	createFailedMessageServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_failed_message/services"
	createRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/events"
	createRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/create_restaurant/services"
	deleteBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/events"
	deleteBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_book/services"
	deleteRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/events"
	deleteRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
	getRestaurantGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/gin_endpoints"
	getRestaurantGRPCService "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/grpc"
	getRestaurantQueries "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/queries"
	listCategoriesGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/gin_endpoints"
	listCategoriesQueries "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/queries"
	restoreRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/events"
	restoreRestaurantServices "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/restore_restaurant/services"
	searchBooksGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_books/gin_endpoints"
	searchBooksQueries "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_books/queries"
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
		createBookServices.NewCreateBookHandler,
		deleteBookServices.NewDeleteBookHandler,
		syncCategoriesServices.NewSyncCategoriesHandler,
		createFailedMessageServices.NewCreateFailedMessageHandler,
	),

	// queries
	fx.Provide(
		getRestaurantQueries.NewGetRestaurantHandler,
		listCategoriesQueries.NewListCategoriesHandler,
		searchBooksQueries.NewSearchBooksHandler,
		searchRestaurantsQueries.NewSearchRestaurantsHandler,
	),

	// events
	fx.Provide(
		createRestaurantEvents.NewCreateRestaurantHandler,
		deleteRestaurantEvents.NewDeleteRestaurantHandler,
		restoreRestaurantEvents.NewRestoreRestaurantHandler,
		updateRestaurantEvents.NewUpdateRestaurantHandler,
		createBookEvents.NewCreateBookHandler,
		deleteBookEvents.NewDeleteBookHandler,
		syncCategoriesEvents.NewSyncCategoriesHandler,
		createFailedMessageEvents.NewCreateFailedMessageHandler,
	),

	// grpc service register
	fx.Provide(
		fxTagEndpoints(
			enums.FxGroupGRPCServiceRegister,
			getRestaurantGRPCService.NewGRPCServiceRegister,
			searchRestaurantsGRPCService.NewGRPCServiceRegister,
		)...,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			enums.FxGroupEndpoints,
			getRestaurantGinEndpoints.NewGetRestaurant,
			searchRestaurantsGinEndpoints.NewSearchRestaurants,
			searchBooksGinEndpoints.NewSearchBooks,
			listCategoriesGinEndpoints.NewListCategories,
		)...,
	),
)

// fxTagEndpoints will tag the endpoints with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagEndpoints(group enums.FxGroup, handlers ...any) []any {
	return lo.Map(handlers, func(handler any, _ int) any {
		return fxTagEndpoint(group, handler)
	})
}

// fxTagEndpoint will tag the endpoint with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagEndpoint(group enums.FxGroup, handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(customizegin.Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, group.ToString())),
	)
}
