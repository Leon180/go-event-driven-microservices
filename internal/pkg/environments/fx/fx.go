package environmentsfx

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/environments"
	"go.uber.org/fx"
)

// environments provide module provide enums.Environment(system environment while running application)
// with no dependencies
var ProvideModule = fx.Module(
	"environmentsProvideFx",
	fx.Provide(func() enums.Environment {
		return environments.InitEnv()
	}),
)
