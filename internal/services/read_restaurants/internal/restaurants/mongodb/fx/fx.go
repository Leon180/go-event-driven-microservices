package mongodbfx

import (
	mongodbfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/mongodb/fx"
	mongodb "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/mongodb"
	"go.uber.org/fx"
)

var ProvideModule = fx.Module(
	"mongodbProvideFx",
	mongodbfx.ProvideModule,
	fx.Provide(
		mongodb.ProvideCollections,
	),
)
