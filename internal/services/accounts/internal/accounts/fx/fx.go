package accountsfx

import (
	customizeginfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/fx"
	environmentsfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/environments/fx"
	loggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/fx"
	contextloggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers/fx"
	uuidfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/fx"
	appconfigsfx "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/configs/fx"
	featuresfx "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/fx"
	postgresdbfx "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/fx"

	"go.uber.org/fx"
)

var AccountsConfiguratorModule = fx.Module(
	"accountsConfiguratorFx",

	ProvideModule,
	InvokeModule,
)

var ProvideModule = fx.Module(
	"accountsProvideFx",

	// environments
	environmentsfx.ProvideModule,

	// uuid generators
	uuidfx.ProvideModule,

	// loggers
	loggersfx.ProvideModule,
	contextloggersfx.ProvideModule,

	// db
	postgresdbfx.ProvideModule,

	// app configs
	appconfigsfx.ProvideModule,

	// features
	featuresfx.ProvideModule,

	// gin server
	customizeginfx.ProvideModule,
)

var InvokeModule = fx.Module(
	"accountsInvokeFx",

	// migrations
	postgresdbfx.InvokeModule,

	// start server
	customizeginfx.InvokeModule,
)
