package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_book/events"
	createBookGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_book/gin_endpoints"
	createBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_book/services"
	deleteBookEvents "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/delete_book/events"
	deleteBookGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/delete_book/gin_endpoints"
	deleteBookServices "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/delete_book/services"
	searchBooksGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/search_books/gin_endpoints"
	searchBooksServices "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/search_books/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the accounts features
// It provides the services and customizeginendpoints.Endpoint for the restaurants features:
// dependencies:
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"restaurantsFeaturesProvideFx",

	// events
	fx.Provide(
		createBookEvents.NewCreateBookMessageBuilder,
		deleteBookEvents.NewDeleteBookMessageBuilder,
	),

	// services
	fx.Provide(
		createBookServices.NewCreateBook,
		deleteBookServices.NewDeleteBook,
		searchBooksServices.NewSearchBooks,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createBookGinEndpoints.NewCreateBook,
			deleteBookGinEndpoints.NewDeleteBook,
			searchBooksGinEndpoints.NewSearchBooks,
		)...,
	),
)

// fxTagEndpoints will tag the endpoints with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagEndpoints(handlers ...any) []any {
	return lo.Map(handlers, func(handler any, _ int) any {
		return fxTagEndpoint(handler)
	})
}

// fxTagEndpoint will tag the endpoint with the group: endpoints for usage of the fx framework
// the group is used to register the endpoint to the router in the gin server
func fxTagEndpoint(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(customizegin.Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupEndpoints.ToString())),
	)
}
