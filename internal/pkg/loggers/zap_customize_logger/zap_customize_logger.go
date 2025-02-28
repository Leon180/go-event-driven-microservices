package zapcustomizelogger

import (
	"log"
	"os"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/samber/lo"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger interface {
	loggers.Logger
	InternalLogger() *zap.Logger
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Sync() error
}

// NewZapLogger create new zap logger
func NewZapLogger(
	env enums.Environment,
) ZapLogger {
	cfg, err := loggers.NewLoggerConfig(env)
	if err != nil {
		log.Println("new zap logger error: config is nil")
		return nil
	}
	zapLogger := &zapLogger{level: cfg.GetLogLevel(), env: env}
	zapLogger.logger, zapLogger.sugarLogger = newLogger(cfg, env)
	return zapLogger
}

func newLogger(cfg *loggers.LoggerConfig, env enums.Environment) (*zap.Logger, *zap.SugaredLogger) {
	if cfg == nil {
		log.Println("error: config is nil")
		return nil, nil
	}
	var encoder zapcore.Encoder
	if env.IsProduction() {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoderCfg := zap.NewDevelopmentEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	options := []zap.Option{zap.AddStacktrace(zap.ErrorLevel)}
	if cfg.GetCallerEnabled() {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(1))
	}

	logger := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(cfg.GetLogLevel())),
		options...,
	)

	if cfg.GetEnableTracing() {
		logger = otelzap.New(logger).Logger
	}

	return logger, logger.Sugar()
}

type zapLogger struct {
	level       zapcore.Level
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
	env         enums.Environment
}

func (l *zapLogger) InternalLogger() *zap.Logger {
	return l.logger
}

func (l *zapLogger) Configure(cfg func(internalLog interface{})) {
	cfg(l.logger)
}

// WithName add logger microservice name
func (l *zapLogger) WithName(name string) {
	l.logger, l.sugarLogger = l.logger.Named(name), l.sugarLogger.Named(name)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *zapLogger) Debugw(msg string, fields loggers.Fields) {
	l.logger.Debug(msg, mapToZapFields(fields)...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Infow logs a message with some additional context.
func (l *zapLogger) Infow(msg string, fields loggers.Fields) {
	l.logger.Info(msg, mapToZapFields(fields)...)
}

// Printf uses fmt.Sprintf to log a templated message
func (l *zapLogger) Printf(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// WarnMsg log error message with warn level.
func (l *zapLogger) WarnMsg(msg string, err error) {
	l.logger.Warn(msg, zap.String("error", err.Error()))
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorw logs a message with some additional context.
func (l *zapLogger) Errorw(msg string, fields loggers.Fields) {
	l.logger.Error(msg, mapToZapFields(fields)...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// Err uses error to log a message.
func (l *zapLogger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *zapLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *zapLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

// Sync flushes any buffered log entries
func (l *zapLogger) Sync() error {
	go func() {
		if err := l.logger.Sync(); err != nil {
			l.logger.Error("error while syncing", zap.Error(err))
		}
	}() // nolint: errcheck
	return l.sugarLogger.Sync()
}

func (l *zapLogger) GetEnvironment() enums.Environment {
	return l.env
}

func (l *zapLogger) GRPCMiddlewareAccessLogger(
	method string,
	time time.Duration,
	metaData map[string][]string,
	err error,
) {
	l.Info(
		enums.GRPC,
		zap.String(enums.METHOD, method),
		zap.Duration(enums.TIME, time),
		zap.Any(enums.METADATA, metaData),
		zap.Error(err),
	)
}

func (l *zapLogger) GRPCClientInterceptorLogger(
	method string,
	req, reply interface{},
	time time.Duration,
	metaData map[string][]string,
	err error,
) {
	l.Info(
		enums.GRPC,
		zap.String(enums.METHOD, method),
		zap.Any(enums.REQUEST, req),
		zap.Any(enums.REPLY, reply),
		zap.Duration(enums.TIME, time),
		zap.Any(enums.METADATA, metaData),
		zap.Error(err),
	)
}

func mapToZapFields(data map[string]interface{}) []zap.Field {
	return lo.MapToSlice(data, func(key string, value interface{}) zap.Field {
		return zap.Field{
			Key:       key,
			Type:      getFieldType(value),
			Interface: value,
		}
	})
}

func getFieldType(value interface{}) zapcore.FieldType {
	switch value.(type) {
	case string:
		return zapcore.StringType
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return zapcore.Int64Type
	case bool:
		return zapcore.BoolType
	case float32, float64:
		return zapcore.Float64Type
	case error:
		return zapcore.ErrorType
	default:
		return zapcore.ReflectType
	}
}
