package utilities

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	zap *zap.Logger
	gormlogger.Config
}

func NewGormZapLogger(zap *zap.Logger) *GormLogger {
	return &GormLogger{
		zap:    zap,
		Config: gormlogger.Config{IgnoreRecordNotFoundError: true},
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.zap.Sugar().Infof(msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.zap.Sugar().Warnf(msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.zap.Sugar().Errorf(msg, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError) {
		l.zap.Error("trace",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
		)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn {
		l.zap.Warn("trace",
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
			zap.String("slow-threshold", l.SlowThreshold.String()),
		)
		return
	}

	if l.LogLevel >= gormlogger.Info {
		l.zap.Info("trace",
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
		)
	}
}
