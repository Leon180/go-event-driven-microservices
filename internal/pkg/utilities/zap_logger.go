package utilities

import (
	"context"
	"os"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	lr          *logRecord
)

type logRecord struct {
	logRecordDay     int
	logFileName      string
	lumberJackLogger *lumberjack.Logger
}

func (l *logRecord) RotateIfNeed() {
	currentDay := time.Now().Day()
	if l.logRecordDay == 0 {
		l.logRecordDay = currentDay
		return
	}
	if l.logRecordDay == currentDay {
		return
	}
	l.logRecordDay = currentDay
	l.lumberJackLogger.Rotate()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger(logCfg configs.LogConfig) (*zap.Logger, *zap.SugaredLogger) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logCfg.FileName,
		MaxSize:    logCfg.MaxSize,
		MaxAge:     logCfg.MaxAge,
		MaxBackups: logCfg.MaxBackups,
		Compress:   logCfg.Compress,
	}
	cfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "severity",
		TimeKey:        "time",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		NameKey:        "logger",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	initLogRecord(0, logCfg.FileName, lumberJackLogger)
	logger = zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg),
				zapcore.NewMultiWriteSyncer(
					zapcore.AddSync(lumberJackLogger),
					zapcore.AddSync(os.Stderr),
				),
				logCfg.Level,
			),
			zapcore.NewCore(
				getEncoder(),
				zapcore.AddSync(os.Stdout),
				logCfg.Level,
			),
		),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.Hooks(func(e zapcore.Entry) error {
			lr.RotateIfNeed()
			return nil
		}),
	)
	sugarLogger = logger.Sugar()
	return logger, sugarLogger
}

func initLogRecord(logRecordDay int, logFileName string, lumberJackLogger *lumberjack.Logger) {
	if lr != nil {
		return
	}
	lr = &logRecord{
		logRecordDay:     logRecordDay,
		logFileName:      logFileName,
		lumberJackLogger: lumberJackLogger,
	}
}

func SyncLogger() {
	logger.Sync()
	sugarLogger.Sync()
}

func getTraceID(ctx context.Context) string {
	value := ctx.Value(enums.TraceIDKey)
	if value == nil {
		return ""
	}
	if traceID, ok := value.(string); ok {
		return traceID
	}
	return ""
}

func LogInfo(ctx context.Context, data ...interface{}) {
	if traceID := getTraceID(ctx); traceID != "" {
		sugarLogger.
			With("EventID", traceID).
			Info(data...)
		return
	}
	sugarLogger.Info(data...)
}

func LogError(ctx context.Context, data ...interface{}) {
	if traceID := getTraceID(ctx); traceID != "" {
		sugarLogger.
			With("EventID", traceID).
			Error(data...)
		return
	}
	sugarLogger.Error(data...)
}

func LogFatal(ctx context.Context, data ...interface{}) {
	if traceID := getTraceID(ctx); traceID != "" {
		sugarLogger.
			With("EventID", traceID).
			Fatal(data...)
		return
	}
	sugarLogger.Fatal(data...)
}
