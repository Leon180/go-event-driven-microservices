package loggersconfigs

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	LogLevel      zapcore.Level `mapstructure:"logLevel"`
	CallerEnabled bool          `mapstructure:"callerEnabled"`
	EnableTracing bool          `mapstructure:"enableTracing"`
}

func (o *Logger) GetLogLevel() zapcore.Level {
	return o.LogLevel
}

func (o *Logger) GetCallerEnabled() bool {
	return o.CallerEnabled
}

func (o *Logger) GetEnableTracing() bool {
	return o.EnableTracing
}

func NewLoggerConfig(env enums.Environment) (*Logger, error) {
	typeName := reflect.GetTypeName[Logger]()
	logger, err := configs.BindConfigByKey[Logger](typeName, env)
	if err != nil {
		return nil, err
	}
	return &logger, nil
}
