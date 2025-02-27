package contextloggers

import (
	"context"
	"fmt"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"go.uber.org/zap"
)

//go:generate mockgen -source=logger.go -destination=./mocks/logger_mock.go -package=mocks

// context logger is a logger that can be used to log messages with context information
type ContextLogger interface {
	WithContextInfo(ctx context.Context, keys ...enums.ContextKey) ContextLogger
	clearContextInfo()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	WarnMsg(msg string, err error)
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Err(msg string, err error)
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Printf(template string, args ...interface{})
}

func NewContextLogger(logger loggers.Logger) ContextLogger {
	return &contextLogger{logger: logger}
}

type contextLogger struct {
	logger      loggers.Logger
	contextInfo map[enums.ContextKey]interface{}
}

func (c *contextLogger) WithContextInfo(ctx context.Context, keys ...enums.ContextKey) ContextLogger {
	for _, key := range keys {
		if value := ctx.Value(key); value != nil {
			c.contextInfo[key] = generateMsgByContextKey(key, value)
		}
	}
	return c
}

func (c *contextLogger) addContextInfoToArgs(args []interface{}) []interface{} {
	if len(c.contextInfo) == 0 {
		return args
	}
	newArgs := make([]interface{}, len(args)+len(c.contextInfo))
	copy(newArgs, args)
	i := len(args)
	for _, value := range c.contextInfo {
		newArgs[i] = value
		i++
	}
	return newArgs
}

func (c *contextLogger) clearContextInfo() {
	c.contextInfo = make(map[enums.ContextKey]interface{})
}

func (l *contextLogger) Configure(cfg func(internalLog interface{})) {
	l.logger.Configure(cfg)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *contextLogger) Debug(args ...interface{}) {
	l.logger.Debug(l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *contextLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Info uses fmt.Sprint to construct and log a message
func (l *contextLogger) Info(args ...interface{}) {
	l.logger.Info(l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *contextLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Printf uses fmt.Sprintf to log a templated message
func (l *contextLogger) Printf(template string, args ...interface{}) {
	l.logger.Printf(template, l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *contextLogger) Warn(args ...interface{}) {
	l.logger.Warn(l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// WarnMsg log error message with warn level.
func (l *contextLogger) WarnMsg(msg string, err error) {
	l.logger.WarnMsg(msg, err)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *contextLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Error uses fmt.Sprint to construct and log a message.
func (l *contextLogger) Error(args ...interface{}) {
	l.logger.Error(l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *contextLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Err uses error to log a message.
func (l *contextLogger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *contextLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *contextLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, l.addContextInfoToArgs(args)...)
	l.clearContextInfo()
}

func generateMsgByContextKey(key enums.ContextKey, contextValue interface{}) string {
	switch key {
	case enums.ContextKeySession, enums.ContextKeyTraceID:
		if contextValue == nil {
			return ""
		}
		switch contextValue.(type) {
		case string:
			return fmt.Sprintf("[%s: %s]", key.ToString(), contextValue)
		case int, int8, int16, int32, int64:
			return fmt.Sprintf("[%s: %d]", key.ToString(), contextValue)
		case uint, uint8, uint16, uint32, uint64:
			return fmt.Sprintf("[%s: %d]", key.ToString(), contextValue)
		case float32, float64:
			return fmt.Sprintf("[%s: %f]", key.ToString(), contextValue)
		case bool:
			return fmt.Sprintf("[%s: %t]", key.ToString(), contextValue)
		default:
			return fmt.Sprintf("[%s: %v]", key.ToString(), contextValue)
		}
	default:
		return ""
	}
}
