package postgresdbfx

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/postgresdb"
	postgresdbmigrations "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/postgresdb/migrations"
	"go.uber.org/fx"
)

// ProvideModule is the module for the postgresdb
// It provides:
// - *gormconfigs.GormDB
// - *gorm.DB
// dependencies:
// - enums.Environment
// - gormlogger.GormCustomLogger
var ProvideModule = fx.Module(
	"customersPostgresdbProvideFx",
	fx.Provide(
		postgresdb.NewGormDBConfig,
		postgresdb.NewGormDB,
	),
)

// InvokeModule is the module for the postgresdb
// It invokes the migrate db function
// dependencies:
// - *gorm.DB
var InvokeModule = fx.Module(
	"customersPostgresdbInvokeFx",
	fx.Invoke(
		postgresdbmigrations.MigrateDB,
	),
)
