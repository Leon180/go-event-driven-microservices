package loggersfx

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	gormcustomizelogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/gorm_customize_logger"
	zapcustomizelogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/zap_customize_logger"
	"go.uber.org/fx"
	gormlogger "gorm.io/gorm/logger"
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
		zapcustomizelogger.NewZapLogger,
		fx.Annotate(
			zapcustomizelogger.NewZapLogger,
			fx.As(new(loggers.Logger)),
		),
		gormcustomizelogger.NewGormCustomLogger,
		fx.Annotate(
			gormcustomizelogger.NewGormCustomLogger,
			fx.As(new(gormlogger.Interface)),
		),
	),
)
