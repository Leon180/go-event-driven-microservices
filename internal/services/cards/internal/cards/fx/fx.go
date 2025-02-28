package cardsfx

import (
	customizeginfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/fx"
	environmentsfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/environments/fx"
	loggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/fx"
	contextloggersfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers/fx"
	uuidfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/fx"
	appconfigsfx "github.com/Leon180/go-event-driven-microservices/internal/services/cards/configs/fx"
	featuresfx "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/fx"
	postgresdbfx "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/postgresdb/fx"
	repositoriesfx "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories/fx"
	cardnumberutilitiesfx "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/utilities/card_number/fx"
	"go.uber.org/fx"
)

var CardsConfiguratorModule = fx.Module(
	"cardsConfiguratorFx",

	ProvideModule,
	InvokeModule,
)

var ProvideModule = fx.Module(
	"cardsProvideFx",

	// environments
	environmentsfx.ProvideModule,

	// card number generator
	cardnumberutilitiesfx.ProvideModule,

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
