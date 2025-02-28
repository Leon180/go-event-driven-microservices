package loggers

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	LogLevel      zapcore.Level `mapstructure:"logLevel"`
	CallerEnabled bool          `mapstructure:"callerEnabled"`
	EnableTracing bool          `mapstructure:"enableTracing"`
}

func (o *LoggerConfig) GetLogLevel() zapcore.Level {
	return o.LogLevel
}

func (o *LoggerConfig) GetCallerEnabled() bool {
	return o.CallerEnabled
}

func (o *LoggerConfig) GetEnableTracing() bool {
	return o.EnableTracing
}

func NewLoggerConfig(env enums.Environment) (*LoggerConfig, error) {
	typeName := reflect.GetTypeName[LoggerConfig]()
	logger, err := configs.BindConfigByKey[LoggerConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	return &logger, nil
}
