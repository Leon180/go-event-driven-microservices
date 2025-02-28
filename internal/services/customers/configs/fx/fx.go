package configsfx

import (
	customersconfigs "github.com/Leon180/go-event-driven-microservices/internal/services/customers/configs"
	"go.uber.org/fx"
)

// ProvideModule is the module for the configs
// It provides the configs:
// - ginconfigs.GinConfig
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"customersConfigsProvideFx",
	fx.Provide(
		customersconfigs.NewAppConfig,
	),
)
