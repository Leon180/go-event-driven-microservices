package gormcustomizelogger

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"

	gormlogger "gorm.io/gorm/logger"
)

type GormCustomLogger interface {
	Error(ctx context.Context, str string, args ...interface{})
	Info(ctx context.Context, str string, args ...interface{})
	LogMode(level gormlogger.LogLevel) gormlogger.Interface
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
	Warn(ctx context.Context, str string, args ...interface{})
	GetEnvironment() enums.Environment
}

func NewGormCustomLogger(logger loggers.Logger) GormCustomLogger {
	return &gormCustomLoggerImpl{
		logger: logger,
		config: gormlogger.Config{LogLevel: gormlogger.Info},
		env:    logger.GetEnvironment(),
	}
}

type gormCustomLoggerImpl struct {
	logger loggers.Logger
	config gormlogger.Config
	env    enums.Environment
}

// LogMode set log mode
func (l *gormCustomLoggerImpl) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.config.LogLevel = level
	return &newlogger
}

// Info prints info
func (l gormCustomLoggerImpl) Info(ctx context.Context, str string, args ...interface{}) {
	if l.config.LogLevel >= gormlogger.Info {
		l.logger.Debugf(str, args...)
	}
}

// Warn prints warn messages
func (l gormCustomLoggerImpl) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.config.LogLevel >= gormlogger.Warn {
		l.logger.Warnf(str, args...)
	}
}

// Error prints error messages
func (l gormCustomLoggerImpl) Error(ctx context.Context, str string, args ...interface{}) {
	if l.config.LogLevel >= gormlogger.Error {
		l.logger.Errorf(str, args...)
	}
}

// Trace prints trace messages
func (l gormCustomLoggerImpl) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (string, int64),
	err error,
) {
	if l.config.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	if l.config.LogLevel >= gormlogger.Info {
		sql, rows := fc()
		l.logger.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.config.LogLevel >= gormlogger.Warn {
		sql, rows := fc()
		l.logger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.config.LogLevel >= gormlogger.Error {
		sql, rows := fc()
		l.logger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}
}

func (l gormCustomLoggerImpl) GetEnvironment() enums.Environment {
	return l.env
}
