package featuresfx

import (
	"fmt"

	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	createBooksEvents "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/events"
	createBooksGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/gin_endpoints"
	createBooksServices "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/services"
	searchBooksGinEndpoints "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/search_books/gin_endpoints"
	searchBooksServices "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/search_books/services"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

// ProvideModule is the module for the books features
// It provides the services and customizeginendpoints.Endpoint for the books features:
// dependencies:
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"restaurantsFeaturesProvideFx",

	// events
	fx.Provide(
		createBooksEvents.NewCreateBookMessageBuilder,
	),

	// services
	fx.Provide(
		createBooksServices.NewCreateBook,
		searchBooksServices.NewSearchBooks,
	),

	// endpoints
	fx.Provide(
		fxTagEndpoints(
			createBooksGinEndpoints.NewCreateBook,
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
