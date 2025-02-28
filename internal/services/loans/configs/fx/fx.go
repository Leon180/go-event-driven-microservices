package configsfx

import (
	cardsconfigs "github.com/Leon180/go-event-driven-microservices/internal/services/cards/configs"
	"go.uber.org/fx"
)

// ProvideModule is the module for the configs
// It provides the configs:
// - ginconfigs.GinConfig
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"cardsConfigsProvideFx",
	fx.Provide(
		cardsconfigs.NewAppConfig,
	),
)
