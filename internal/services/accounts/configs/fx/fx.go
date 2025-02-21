package configsfx

import (
	ginconfigs "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/configs"
	"go.uber.org/fx"
)

// ProvideModule is the module for the configs
// It provides the configs:
// - *configs.App
// - ginconfigs.GinConfig
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"configsProvideFx",
	fx.Provide(
		configs.NewAppConfig,
		fx.Annotate(
			configs.NewAppConfig,
			fx.As(new(ginconfigs.GinConfig)),
		),
	),
)
