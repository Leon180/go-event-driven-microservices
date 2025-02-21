package uuidfx

import (
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"go.uber.org/fx"
)

// uuid provide module provide uuid generator
// with no dependencies
var ProvideModule = fx.Module(
	"uuidProvideFx",
	fx.Provide(
		uuid.NewUUIDGenerator,
	),
)
