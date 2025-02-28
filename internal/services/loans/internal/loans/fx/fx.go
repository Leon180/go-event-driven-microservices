package cardsfx

import (
	customizeginfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/fx"
	environmentsfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/environments/fx"
	loggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/fx"
	contextloggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers/fx"
	uuidfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/fx"
	appconfigsfx "github.com/Leon180/go-event-driven-microservices/internal/services/loans/configs/fx"
	featuresfx "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/fx"
	postgresdbfx "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/postgresdb/fx"
	repositoriesfx "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories/fx"
	loannumberutilitiesfx "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/utilities/loan_number/fx"
	"go.uber.org/fx"
)

var LoansConfiguratorModule = fx.Module(
	"loansConfiguratorFx",

	ProvideModule,
	InvokeModule,
)

var ProvideModule = fx.Module(
	"cardsProvideFx",

	// environments
	environmentsfx.ProvideModule,

	// loan number generator
	loannumberutilitiesfx.ProvideModule,

	// uuid generators
	uuidfx.ProvideModule,

	// loggers
	loggersfx.ProvideModule,
	contextloggersfx.ProvideModule,

	// db
	postgresdbfx.ProvideModule,

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
	"cardsInvokeFx",

	// migrations
	postgresdbfx.InvokeModule,

	// start server
	customizeginfx.InvokeModule,
)
