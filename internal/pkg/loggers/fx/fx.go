package loggersfx

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	gormlogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/grom_logger"
	zaplogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/zap_logger"
	"go.uber.org/fx"
	gormiologger "gorm.io/gorm/logger"
)

// loggers provide module provide loggers:
// - zaplogger.Logger
// - gormlogger.GormCustomLogger
// - loggers.Logger
// - gormiologger.Interface
// dependencies:
// - enums.Environment
var ProvideModule = fx.Module(
	"loggersProvideFx",
	fx.Provide(
		zaplogger.NewZapLogger,
		fx.Annotate(
			zaplogger.NewZapLogger,
			fx.As(new(loggers.Logger)),
		),
		gormlogger.NewGormCustomLogger,
		fx.Annotate(
			gormlogger.NewGormCustomLogger,
			fx.As(new(gormiologger.Interface)),
		),
	),
)
