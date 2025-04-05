package configsfx

import (
	booksconfigs "github.com/Leon180/go-event-driven-microservices/internal/services/books/configs"
	"go.uber.org/fx"
)

// ProvideModule is the module for the configs
// It provides the configs:
// - ginconfigs.GinConfig
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"configsProvideFx",
	fx.Provide(
		booksconfigs.NewAppConfig,
	),
)
