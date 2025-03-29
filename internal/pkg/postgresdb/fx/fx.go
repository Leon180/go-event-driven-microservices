package postgresdbfx

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/postgresdb"
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
	"postgresdbProvideFx",
	fx.Provide(
		postgresdb.NewGormDBConfig,
		postgresdb.NewGormDB,
	),
)
