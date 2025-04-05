package accountsfx

import (
	customizeginfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/fx"
	environmentsfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/environments/fx"
	loggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/fx"
	contextloggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers/fx"
	uuidfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/fx"
	appconfigsfx "github.com/Leon180/go-event-driven-microservices/internal/services/books/configs/fx"
	featuresfx "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/fx"
	repositoriesfx "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories/fx"
	"go.uber.org/fx"
)

var BooksConfiguratorModule = fx.Module(
	"booksConfiguratorFx",

	ProvideModule,
	InvokeModule,
)

var ProvideModule = fx.Module(
	"booksProvideFx",

	// environments
	environmentsfx.ProvideModule,

	// uuid generators
	uuidfx.ProvideModule,

	// loggers
	loggersfx.ProvideModule,
	contextloggersfx.ProvideModule,

	// app configs
	appconfigsfx.ProvideModule,

	// repositories
	repositoriesfx.ProvideModule,

	// features
	featuresfx.ProvideModule,

	// gin server
	customizeginfx.ProvideModule,
)

var InvokeModule = fx.Module(
	"booksInvokeFx",

	// start server
	customizeginfx.InvokeModule,
)
