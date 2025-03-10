package postgresdbfx

import (
	postgresdbfx "github.com/Leon180/go-event-driven-microservices/internal/pkg/postgresdb/fx"
	postgresdbmigrations "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/postgresdb/migrations"
	"go.uber.org/fx"
)

// ProvideModule is the module for the postgresdb
// It provides:
// - *gormconfigs.GormDB
// - *gorm.DB
// dependencies:
// - enums.Environment
// - gormlogger.GormCustomLogger
var ProvideModule = postgresdbfx.ProvideModule

// InvokeModule is the module for the postgresdb
// It invokes the migrate db function
// dependencies:
// - *gorm.DB
var InvokeModule = fx.Module(
	"postgresdbInvokeFx",
	fx.Invoke(
		postgresdbmigrations.MigrateDB,
	),
)
